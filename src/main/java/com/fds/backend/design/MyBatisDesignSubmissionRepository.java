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
    public void create(UUID id, String fullName, String address, String phone, String note, String imageUrl) {
        mapper.create(id, fullName, address, phone, note, imageUrl);
    }

    @Override
    public List<DesignSubmission> list(String status, String keyword, int limit, int offset) {
        return mapper.list(status, keyword, limit, offset);
    }

    @Override
    public void updateStatus(UUID id, String status) {
        mapper.updateStatus(id, status);
    }

    @Override
    public void delete(UUID id) {
        mapper.delete(id);
    }
}
