package com.fds.backend.file;

import org.springframework.web.multipart.MultipartFile;

import java.io.IOException;

public interface FileStorage {
    String save(MultipartFile file) throws IOException;
    void delete(String relativePath) throws IOException;
}
