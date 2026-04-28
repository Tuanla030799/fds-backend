package com.fds.backend.design;

import com.fds.backend.auth.CurrentAdmin;
import com.fds.backend.common.ApiException;
import com.fds.backend.file.FileAsset;
import com.fds.backend.file.FileAssetRepository;
import com.fds.backend.file.FileStorage;
import com.fds.backend.file.FileUrlService;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.io.IOException;
import java.util.List;
import java.util.UUID;

@Service
public class DesignSubmissionService {
    private final DesignSubmissionRepository repository;
    private final FileAssetRepository fileAssetRepository;
    private final FileStorage fileStorage;
    private final FileUrlService fileUrlService;
    private final CurrentAdmin currentAdmin;

    public DesignSubmissionService(DesignSubmissionRepository repository, FileAssetRepository fileAssetRepository,
                                   FileStorage fileStorage, FileUrlService fileUrlService, CurrentAdmin currentAdmin) {
        this.repository = repository;
        this.fileAssetRepository = fileAssetRepository;
        this.fileStorage = fileStorage;
        this.fileUrlService = fileUrlService;
        this.currentAdmin = currentAdmin;
    }

    @Transactional
    public void create(String fullName, String address, String phone, String note, UUID fileId) {
        var file = fileAssetRepository.findById(fileId);
        if (file == null) {
            throw new ApiException("Design submission image file not found");
        }
        var adminId = currentAdmin.idOrNull();
        repository.create(UUID.randomUUID(), fullName, address, phone, note, file.id(), adminId);
        fileAssetRepository.updateStatus(file.id(), "ACTIVE", adminId);
    }

    public List<DesignSubmission> list(String status, String keyword, int page, int limit) {
        int safeLimit = Math.min(Math.max(limit, 1), 100);
        int offset = (Math.max(page, 1) - 1) * safeLimit;
        return repository.list(status, keyword, safeLimit, offset).stream()
                .map(this::withImageUrl)
                .toList();
    }

    @Transactional
    public void updateStatus(UUID id, String status) {
        repository.updateStatus(id, status, currentAdmin.idOrNull());
    }

    @Transactional
    public void delete(UUID id) {
        var submission = repository.findById(id);
        if (submission == null) {
            throw new ApiException("Design submission not found");
        }
        repository.delete(id, currentAdmin.idOrNull());
        deleteImage(submission.fileId());
    }

    private DesignSubmission withImageUrl(DesignSubmission submission) {
        if (submission == null || submission.imageUrl() == null || submission.imageUrl().isBlank()) {
            return submission;
        }
        return new DesignSubmission(
                submission.id(),
                submission.fullName(),
                submission.address(),
                submission.phone(),
                submission.note(),
                submission.fileId(),
                fileUrlService.publicUrl(submission.imageUrl()),
                submission.status(),
                submission.createdAt()
        );
    }

    private void deleteImage(UUID fileId) {
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
            throw new IllegalStateException("Failed to delete design submission image file", e);
        }
        fileAssetRepository.deleteById(file.id());
    }
}
