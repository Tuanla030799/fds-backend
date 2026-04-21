package com.fds.backend.file;

import com.fds.backend.common.ApiResponse;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.MediaType;
import org.springframework.util.StringUtils;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.multipart.MultipartFile;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.Map;
import java.util.Set;
import java.util.UUID;

@RestController
@RequestMapping("/api/files")
public class FileController {
    private final FileAssetMapper fileAssetMapper;
    private final Path uploadRoot;
    private static final Set<String> ALLOWED_FOLDERS = Set.of("tmp", "presets", "design-submissions");

    public FileController(FileAssetMapper fileAssetMapper, @Value("${app.upload-dir}") String uploadDir) {
        this.fileAssetMapper = fileAssetMapper;
        this.uploadRoot = Path.of(uploadDir);
    }

    @PostMapping(value = "/upload", consumes = MediaType.MULTIPART_FORM_DATA_VALUE)
    public ApiResponse<Map<String, String>> upload(@RequestParam("file") MultipartFile file,
                                                    @RequestParam("folder") String folder) throws IOException {
        if (!ALLOWED_FOLDERS.contains(folder)) {
            throw new IllegalArgumentException("Invalid folder");
        }
        Files.createDirectories(uploadRoot.resolve(folder));
        String clean = StringUtils.cleanPath(file.getOriginalFilename());
        String name = UUID.randomUUID() + "-" + clean;
        Path target = uploadRoot.resolve(folder).resolve(name);
        Files.write(target, file.getBytes());

        UUID id = UUID.randomUUID();
        String relativePath = folder + "/" + name;
        fileAssetMapper.create(id, relativePath, folder, "UNACTIVE");
        return ApiResponse.ok("Uploaded", Map.of("fileId", id.toString(), "path", relativePath));
    }
}
