package com.fds.backend.preset;

import com.fds.backend.auth.CurrentAdmin;
import com.fds.backend.file.FileAsset;
import com.fds.backend.file.FileAssetRepository;
import org.junit.jupiter.api.Test;
import org.mockito.ArgumentCaptor;

import java.util.UUID;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.mockito.Mockito.*;

class PresetServiceTest {

    @Test
    void createShouldApplyDefaults() {
        PresetRepository repository = mock(PresetRepository.class);
        FileAssetRepository fileAssetRepository = mock(FileAssetRepository.class);
        CurrentAdmin currentAdmin = mock(CurrentAdmin.class);
        PresetService service = new PresetService(repository, fileAssetRepository, currentAdmin);
        UUID fileId = UUID.randomUUID();
        when(fileAssetRepository.findById(fileId))
                .thenReturn(new FileAsset(fileId, "a.png", "INACTIVE", null));

        service.create(new PresetService.PresetCreateRequest("A", null, null, null, fileId, null));

        ArgumentCaptor<String> statusCap = ArgumentCaptor.forClass(String.class);
        ArgumentCaptor<String> tagsCap = ArgumentCaptor.forClass(String.class);
        ArgumentCaptor<String> imageUrlCap = ArgumentCaptor.forClass(String.class);
        ArgumentCaptor<Integer> sortCap = ArgumentCaptor.forClass(Integer.class);

        verify(repository).create(any(), any(), statusCap.capture(), any(), tagsCap.capture(),
                imageUrlCap.capture(), sortCap.capture(), isNull());
        verify(fileAssetRepository).updateStatus(fileId, "ACTIVE", null);

        assertEquals("active", statusCap.getValue());
        assertEquals("[]", tagsCap.getValue());
        assertEquals("a.png", imageUrlCap.getValue());
        assertEquals(0, sortCap.getValue());
    }
}
