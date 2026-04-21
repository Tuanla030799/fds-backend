package com.fds.backend.file;

import org.springframework.stereotype.Repository;

import java.util.UUID;

@Repository
public class MyBatisFileAssetRepository implements FileAssetRepository {
    private final FileAssetMapper mapper;

    public MyBatisFileAssetRepository(FileAssetMapper mapper) {
        this.mapper = mapper;
    }

    @Override
    public void create(UUID id, String path, String folder, String status) {
        mapper.create(id, path, folder, status);
    }
}
