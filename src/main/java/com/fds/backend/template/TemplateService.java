package com.fds.backend.template;

import com.fds.backend.auth.CurrentAdmin;
import com.fds.backend.common.ApiException;
import com.fds.backend.file.FileAsset;
import com.fds.backend.file.FileAssetRepository;
import com.fds.backend.file.FileUrlService;
import com.fds.backend.file.FileStorage;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.io.IOException;
import java.util.List;
import java.util.UUID;

@Service
public class TemplateService {
    public static final String STATUS_ACTIVE = "ACTIVE";
    public static final String STATUS_INACTIVE = "INACTIVE";

    private final TemplateRepository repository;
    private final FileAssetRepository fileAssetRepository;
    private final FileUrlService fileUrlService;
    private final FileStorage fileStorage;
    private final CurrentAdmin currentAdmin;

    public TemplateService(TemplateRepository repository, FileAssetRepository fileAssetRepository, FileStorage fileStorage,
                           FileUrlService fileUrlService, CurrentAdmin currentAdmin) {
        this.repository = repository;
        this.fileAssetRepository = fileAssetRepository;
        this.fileStorage = fileStorage;
        this.fileUrlService = fileUrlService;
        this.currentAdmin = currentAdmin;
    }

    public List<Template> listPublic(String keyword, int page, int limit) {
        return listInternal(STATUS_ACTIVE, keyword, page, limit);
    }

    public List<Template> listAdmin(String status, String keyword, int page, int limit) {
        return listInternal(normalizeStatusNullable(status), keyword, page, limit);
    }

    private List<Template> listInternal(String status, String keyword, int page, int limit) {
        int safeLimit = Math.min(Math.max(limit, 1), 100);
        int offset = (Math.max(page, 1) - 1) * safeLimit;
        return repository.list(status, keyword, safeLimit, offset).stream()
                .map(this::withImageUrl)
                .toList();
    }

    @Transactional
    public void create(TemplateCreateRequest req) {
        String finalStatus = normalizeStatusNullable(req.status());
        if (finalStatus == null) {
            finalStatus = STATUS_ACTIVE;
        }
        var file = fileAssetRepository.findById(req.fileId());
        if (file == null) {
            throw new ApiException("Template image file not found");
        }

        var adminId = currentAdmin.idOrNull();
        repository.create(UUID.randomUUID(), req.name(), finalStatus, req.note(), file.id(), adminId);
        fileAssetRepository.updateStatus(file.id(), STATUS_ACTIVE, adminId);
    }

    @Transactional
    public void updateStatus(UUID id, String status) {
        if (repository.findById(id) == null) {
            throw new ApiException("Template not found");
        }
        repository.updateStatus(id, requireStatus(status), currentAdmin.idOrNull());
    }

    @Transactional
    public void delete(UUID id) {
        var template = repository.findById(id);
        if (template == null) {
            throw new ApiException("Template not found");
        }

        repository.delete(id, currentAdmin.idOrNull());
        deleteTemplateImage(template.fileId());
    }

    public record TemplateCreateRequest(String name, String status, String note, UUID fileId) {}

    private Template withImageUrl(Template template) {
        if (template == null) {
            return null;
        }
        if (template.imageUrl() == null || template.imageUrl().isBlank()) {
            return template;
        }
        return new Template(
                template.id(),
                template.name(),
                template.status(),
                template.note(),
                template.fileId(),
                fileUrlService.publicUrl(template.imageUrl()),
                template.createdAt()
        );
    }

    private void deleteTemplateImage(UUID fileId) {
        if (fileId == null) {
            return;
        }

        FileAsset file = fileAssetRepository.findById(fileId);
        if (file == null) {
            return;
        }

        try {
            fileStorage.delete(file.path());
        } catch (IOException e) {
            throw new IllegalStateException("Failed to delete template image file", e);
        }

        fileAssetRepository.deleteById(file.id());
    }

    private String normalizeStatusNullable(String status) {
        if (status == null || status.isBlank()) {
            return null;
        }
        String normalized = status.trim().toUpperCase();
        if (!STATUS_ACTIVE.equals(normalized) && !STATUS_INACTIVE.equals(normalized)) {
            throw new ApiException("Template status must be ACTIVE or INACTIVE");
        }
        return normalized;
    }

    private String requireStatus(String status) {
        String normalized = normalizeStatusNullable(status);
        if (normalized == null) {
            throw new ApiException("Template status is required");
        }
        return normalized;
    }
}
