package com.fds.backend.file;

import com.fds.backend.auth.CurrentAdmin;
import com.fds.backend.common.ApiException;
import org.junit.jupiter.api.Test;
import org.springframework.mock.web.MockMultipartFile;

import java.time.OffsetDateTime;
import java.util.List;
import java.util.Map;
import java.util.UUID;

import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.mockito.Mockito.*;

class FileServiceTest {

    @Test
    void uploadShouldRejectEmptyFile() {
        FileAssetRepository repository = mock(FileAssetRepository.class);
        FileStorage storage = mock(FileStorage.class);
        CurrentAdmin currentAdmin = mock(CurrentAdmin.class);
        FileUrlService fileUrlService = new FileUrlService("/files");
        FileService service = new FileService(repository, storage, currentAdmin, fileUrlService);

        MockMultipartFile file = new MockMultipartFile("file", "a.png", "image/png", new byte[]{});

        assertThrows(ApiException.class, () -> service.upload(file));
    }

    @Test
    void uploadShouldCreateInactiveFileRecord() throws Exception {
        FileAssetRepository repository = mock(FileAssetRepository.class);
        FileStorage storage = mock(FileStorage.class);
        CurrentAdmin currentAdmin = mock(CurrentAdmin.class);
        FileUrlService fileUrlService = new FileUrlService("/files");
        FileService service = new FileService(repository, storage, currentAdmin, fileUrlService);
        MockMultipartFile file = new MockMultipartFile("file", "a.png", "image/png", new byte[]{1, 2, 3});
        when(storage.save(file)).thenReturn("a.png");

        Map<String, String> result = service.upload(file);

        verify(repository).create(any(UUID.class), eq("a.png"), eq("INACTIVE"), isNull());
        assertEquals("/files/a.png", result.get("url"));
    }

    @Test
    void uploadShouldUseConfiguredPublicBaseUrl() throws Exception {
        FileAssetRepository repository = mock(FileAssetRepository.class);
        FileStorage storage = mock(FileStorage.class);
        CurrentAdmin currentAdmin = mock(CurrentAdmin.class);
        FileUrlService fileUrlService = new FileUrlService("https://cdn.example.com/files/");
        FileService service = new FileService(repository, storage, currentAdmin, fileUrlService);
        MockMultipartFile file = new MockMultipartFile("file", "a.png", "image/png", new byte[]{1, 2, 3});
        when(storage.save(file)).thenReturn("2026-01-01/a.png");

        Map<String, String> result = service.upload(file);

        assertEquals("https://cdn.example.com/files/2026-01-01/a.png", result.get("url"));
    }

    @Test
    void cleanupInactiveFilesShouldDeleteStorageAndRecord() throws Exception {
        FileAssetRepository repository = mock(FileAssetRepository.class);
        FileStorage storage = mock(FileStorage.class);
        CurrentAdmin currentAdmin = mock(CurrentAdmin.class);
        FileUrlService fileUrlService = new FileUrlService("/files");
        FileService service = new FileService(repository, storage, currentAdmin, fileUrlService);
        UUID fileId = UUID.randomUUID();
        OffsetDateTime cutoff = OffsetDateTime.now().minusMinutes(10);
        when(repository.listInactiveCreatedBefore(cutoff))
                .thenReturn(List.of(new FileAsset(fileId, "a.png", "INACTIVE", null)));

        service.cleanupInactiveFiles(cutoff);

        verify(storage).delete("a.png");
        verify(repository).deleteById(fileId);
    }
}
