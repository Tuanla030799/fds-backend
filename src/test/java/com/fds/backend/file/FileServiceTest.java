package com.fds.backend.file;

import com.fds.backend.common.ApiException;
import org.junit.jupiter.api.Test;
import org.springframework.mock.web.MockMultipartFile;

import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.mockito.Mockito.mock;

class FileServiceTest {

    @Test
    void uploadShouldRejectInvalidFolder() {
        FileAssetRepository repository = mock(FileAssetRepository.class);
        FileStorage storage = mock(FileStorage.class);
        FileService service = new FileService(repository, storage);

        MockMultipartFile file = new MockMultipartFile("file", "a.png", "image/png", new byte[]{1, 2, 3});

        assertThrows(ApiException.class, () -> service.upload(file, "invalid-folder"));
    }
}
