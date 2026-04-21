package com.fds.backend.file;

import com.fds.backend.common.ApiException;
import org.springframework.stereotype.Service;
import org.springframework.web.multipart.MultipartFile;

import java.io.IOException;
import java.util.Map;
import java.util.Set;
import java.util.UUID;

@Service
public class FileService {
    private static final Set<String> ALLOWED_FOLDERS = Set.of("tmp", "presets", "design-submissions");

    private final FileAssetRepository repository;
    private final FileStorage storage;

    public FileService(FileAssetRepository repository, FileStorage storage) {
        this.repository = repository;
        this.storage = storage;
    }

    public Map<String, String> upload(MultipartFile file, String folder) throws IOException {
        if (!ALLOWED_FOLDERS.contains(folder)) {
            throw new ApiException("Invalid folder");
        }
        String relativePath = storage.save(file, folder);
        UUID id = UUID.randomUUID();
        repository.create(id, relativePath, folder, "UNACTIVE");
        return Map.of("fileId", id.toString(), "path", relativePath);
    }
}
