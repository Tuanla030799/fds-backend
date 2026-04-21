package com.fds.backend.file;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;
import org.springframework.util.StringUtils;
import org.springframework.web.multipart.MultipartFile;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.UUID;

@Component
public class LocalFileStorage implements FileStorage {
    private final Path uploadRoot;

    public LocalFileStorage(@Value("${app.upload-dir}") String uploadDir) {
        this.uploadRoot = Path.of(uploadDir);
    }

    @Override
    public String save(MultipartFile file, String folder) throws IOException {
        Files.createDirectories(uploadRoot.resolve(folder));
        String clean = StringUtils.cleanPath(file.getOriginalFilename());
        String name = UUID.randomUUID() + "-" + clean;
        Path target = uploadRoot.resolve(folder).resolve(name);
        Files.write(target, file.getBytes());
        return folder + "/" + name;
    }
}
