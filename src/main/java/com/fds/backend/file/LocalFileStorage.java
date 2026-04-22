package com.fds.backend.file;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;
import org.springframework.util.StringUtils;
import org.springframework.web.multipart.MultipartFile;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.time.LocalDate;
import java.util.UUID;

@Component
public class LocalFileStorage implements FileStorage {
    private static final Logger log = LoggerFactory.getLogger(LocalFileStorage.class);

    private final Path uploadRoot;

    public LocalFileStorage(@Value("${app.upload-dir}") String uploadDir) {
        this.uploadRoot = Path.of(uploadDir);
    }

    @Override
    public String save(MultipartFile file) throws IOException {
        String dateFolder = LocalDate.now().toString();
        Path targetDir = uploadRoot.resolve(dateFolder);
        Files.createDirectories(targetDir);
        String clean = StringUtils.cleanPath(file.getOriginalFilename());
        String name = UUID.randomUUID() + "-" + clean;
        Path target = targetDir.resolve(name);
        Files.write(target, file.getBytes());
        log.info("File stored path={} size={}", target.toAbsolutePath(), file.getSize());
        return dateFolder + "/" + name;
    }

    @Override
    public void delete(String relativePath) throws IOException {
        Path target = uploadRoot.resolve(relativePath).normalize();
        if (!target.startsWith(uploadRoot.normalize())) {
            return;
        }
        Files.deleteIfExists(target);
    }
}
