package com.fds.backend.design;

import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.UUID;

@Repository
public class MyBatisDesignSubmissionRepository implements DesignSubmissionRepository {
    private final DesignSubmissionMapper mapper;

    public MyBatisDesignSubmissionRepository(DesignSubmissionMapper mapper) {
        this.mapper = mapper;
    }

    @Override
    public void create(UUID id, String fullName, String address, String phone, String note, String imageUrl,
                       UUID createdBy) {
        mapper.create(id, fullName, address, phone, note, imageUrl, createdBy);
    }

    @Override
    public List<DesignSubmission> list(String status, String keyword, int limit, int offset) {
        return mapper.list(status, keyword, limit, offset);
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
