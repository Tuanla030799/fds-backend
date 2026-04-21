package com.fds.backend.design;

import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;
import java.util.UUID;

@Service
public class DesignSubmissionService {
    private final DesignSubmissionRepository repository;

    public DesignSubmissionService(DesignSubmissionRepository repository) {
        this.repository = repository;
    }

    @Transactional
    public void create(String fullName, String address, String phone, String note, String imageUrl) {
        repository.create(UUID.randomUUID(), fullName, address, phone, note, imageUrl);
    }

    public List<DesignSubmission> list(String status, String keyword, int page, int limit) {
        int safeLimit = Math.min(Math.max(limit, 1), 100);
        int offset = (Math.max(page, 1) - 1) * safeLimit;
        return repository.list(status, keyword, safeLimit, offset);
    }

    @Transactional
    public void updateStatus(UUID id, String status) {
        repository.updateStatus(id, status);
    }

    @Transactional
    public void delete(UUID id) {
        repository.delete(id);
    }
}
