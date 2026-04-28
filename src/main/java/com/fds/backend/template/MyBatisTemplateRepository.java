package com.fds.backend.template;

import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.UUID;

@Repository
public class MyBatisTemplateRepository implements TemplateRepository {
    private final TemplateMapper mapper;

    public MyBatisTemplateRepository(TemplateMapper mapper) {
        this.mapper = mapper;
    }

    @Override
    public List<Template> list(String status, String keyword, int limit, int offset) {
        return mapper.list(status, keyword, limit, offset);
    }

    @Override
    public Template findById(UUID id) {
        return mapper.findById(id);
    }

    @Override
    public void create(UUID id, String name, String status, String note, UUID fileId, UUID createdBy) {
        mapper.create(id, name, status, note, fileId, createdBy);
    }

    @Override
    public void updateStatus(UUID id, String status, UUID updatedBy) {
        mapper.updateStatus(id, status, updatedBy);
    }

    @Override
    public void delete(UUID id, UUID updatedBy) {
        mapper.delete(id, updatedBy);
    }
}
