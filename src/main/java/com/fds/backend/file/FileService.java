package com.fds.backend.file;

import com.fds.backend.auth.CurrentAdmin;
import com.fds.backend.common.ApiException;
import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.web.multipart.MultipartFile;

import java.io.IOException;
import java.time.OffsetDateTime;
import java.util.Map;
import java.util.UUID;

@Service
public class FileService {
    private static final String STATUS_INACTIVE = "INACTIVE";

    private final FileAssetRepository repository;
    private final FileStorage storage;
    private final CurrentAdmin currentAdmin;
    private final FileUrlService fileUrlService;

    public FileService(FileAssetRepository repository, FileStorage storage, CurrentAdmin currentAdmin,
                       FileUrlService fileUrlService) {
        this.repository = repository;
        this.storage = storage;
        this.currentAdmin = currentAdmin;
        this.fileUrlService = fileUrlService;
    }

    @Transactional
    public Map<String, String> upload(MultipartFile file) throws IOException {
        if (file.isEmpty()) {
            throw new ApiException("File is empty");
        }
        String relativePath = storage.save(file);
        UUID id = UUID.randomUUID();
        repository.create(id, relativePath, STATUS_INACTIVE, currentAdmin.idOrNull());
        return Map.of("fileId", id.toString(), "path", relativePath, "url", fileUrlService.publicUrl(relativePath));
    }

    @Scheduled(fixedDelayString = "${app.file-cleanup.fixed-delay-ms:60000}")
    public void cleanupInactiveFiles() {
        cleanupInactiveFiles(OffsetDateTime.now().minusMinutes(10));
    }

    @Transactional
    public void cleanupInactiveFiles(OffsetDateTime cutoff) {
        var files = repository.listInactiveCreatedBefore(cutoff);
        for (var file : files) {
            try {
                storage.delete(file.path());
            } catch (IOException ignored) {
                // Keep DB cleanup moving; missing local files should not block stale record removal.
            }
            repository.deleteById(file.id());
        }
    }
}
