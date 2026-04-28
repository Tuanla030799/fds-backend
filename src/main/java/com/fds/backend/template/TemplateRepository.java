package com.fds.backend.template;

import java.util.List;
import java.util.UUID;

public interface TemplateRepository {
    List<Template> list(String status, String keyword, int limit, int offset);
    Template findById(UUID id);
    void create(UUID id, String name, String status, String note, UUID fileId, UUID createdBy);
    void updateStatus(UUID id, String status, UUID updatedBy);
    void delete(UUID id, UUID updatedBy);
}
