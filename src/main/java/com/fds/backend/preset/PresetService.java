package com.fds.backend.preset;

import com.fds.backend.auth.CurrentAdmin;
import com.fds.backend.common.ApiException;
import com.fds.backend.file.FileAssetRepository;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;
import java.util.UUID;

@Service
public class PresetService {
    private final PresetRepository repository;
    private final FileAssetRepository fileAssetRepository;
    private final CurrentAdmin currentAdmin;

    public PresetService(PresetRepository repository, FileAssetRepository fileAssetRepository, CurrentAdmin currentAdmin) {
        this.repository = repository;
        this.fileAssetRepository = fileAssetRepository;
        this.currentAdmin = currentAdmin;
    }

    public List<Preset> list(String status, String keyword, int page, int limit) {
        int safeLimit = Math.min(Math.max(limit, 1), 100);
        int offset = (Math.max(page, 1) - 1) * safeLimit;
        return repository.list(status, keyword, safeLimit, offset);
    }

    @Transactional
    public void create(PresetCreateRequest req) {
        String finalStatus = (req.status() == null || req.status().isBlank()) ? "active" : req.status();
        String finalTags = (req.tags() == null || req.tags().isBlank()) ? "[]" : req.tags();
        int finalSort = req.sortOrder() == null ? 0 : req.sortOrder();
        var file = fileAssetRepository.findById(req.fileId());
        if (file == null) {
            throw new ApiException("Preset image file not found");
        }

        var adminId = currentAdmin.idOrNull();
        repository.create(UUID.randomUUID(), req.name(), finalStatus, req.note(), finalTags, file.path(), finalSort,
                adminId);
        fileAssetRepository.updateStatus(file.id(), "ACTIVE", adminId);
    }

    @Transactional
    public void delete(UUID id) {
        repository.delete(id, currentAdmin.idOrNull());
    }

    public record PresetCreateRequest(String name, String status, String note, String tags, UUID fileId, Integer sortOrder) {}
}
