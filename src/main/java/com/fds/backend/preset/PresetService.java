package com.fds.backend.preset;

import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;
import java.util.UUID;

@Service
public class PresetService {
    private final PresetRepository repository;

    public PresetService(PresetRepository repository) {
        this.repository = repository;
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

        repository.create(UUID.randomUUID(), req.name(), finalStatus, req.note(), finalTags, req.imageUrl(), finalSort);
    }

    @Transactional
    public void delete(UUID id) {
        repository.delete(id);
    }

    public record PresetCreateRequest(String name, String status, String note, String tags, String imageUrl, Integer sortOrder) {}
}
