package com.fds.backend.preset;

import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.UUID;

@Repository
public class MyBatisPresetRepository implements PresetRepository {
    private final PresetMapper mapper;

    public MyBatisPresetRepository(PresetMapper mapper) {
        this.mapper = mapper;
    }

    @Override
    public List<Preset> list(String status, String keyword, int limit, int offset) {
        return mapper.list(status, keyword, limit, offset);
    }

    @Override
    public void create(UUID id, String name, String status, String note, String tags, String imageUrl, int sortOrder,
                       UUID createdBy) {
        mapper.create(id, name, status, note, tags, imageUrl, sortOrder, createdBy);
    }

    @Override
    public void delete(UUID id, UUID updatedBy) {
        mapper.delete(id, updatedBy);
    }
}
