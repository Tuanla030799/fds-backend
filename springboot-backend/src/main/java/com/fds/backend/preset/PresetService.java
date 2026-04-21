package com.fds.backend.preset;

import org.springframework.stereotype.Service;

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

    public void create(PresetCreateRequest req) {
        repository.create(UUID.randomUUID(), req.name(), req.status(), req.note(), req.tags(), req.imageUrl(), req.sortOrder());
    }

    public void delete(UUID id) {
        repository.delete(id);
    }

    public record PresetCreateRequest(String name, String status, String note, String tags, String imageUrl, int sortOrder) {}
}
