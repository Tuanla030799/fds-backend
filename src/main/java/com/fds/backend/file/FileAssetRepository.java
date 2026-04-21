package com.fds.backend.file;

import java.util.UUID;

public interface FileAssetRepository {
    void create(UUID id, String path, String folder, String status);
}
