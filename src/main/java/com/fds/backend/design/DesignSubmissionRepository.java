package com.fds.backend.design;

import java.util.List;
import java.util.UUID;

public interface DesignSubmissionRepository {
    void create(UUID id, String fullName, String address, String phone, String note, UUID fileId, UUID createdBy);
    DesignSubmission findById(UUID id);
    List<DesignSubmission> list(String status, String keyword, int limit, int offset);
    void updateStatus(UUID id, String status, UUID updatedBy);
    void delete(UUID id, UUID updatedBy);
}
