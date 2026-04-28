package com.fds.backend.template;

import com.fds.backend.auth.CurrentAdmin;
import com.fds.backend.file.FileAsset;
import com.fds.backend.file.FileAssetRepository;
import com.fds.backend.file.FileUrlService;
import com.fds.backend.file.FileStorage;
import org.junit.jupiter.api.Test;
import org.mockito.ArgumentCaptor;

import java.util.UUID;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.mockito.Mockito.*;

class TemplateServiceTest {

    @Test
    void createShouldApplyDefaults() {
        TemplateRepository repository = mock(TemplateRepository.class);
        FileAssetRepository fileAssetRepository = mock(FileAssetRepository.class);
        FileStorage fileStorage = mock(FileStorage.class);
        FileUrlService fileUrlService = mock(FileUrlService.class);
        CurrentAdmin currentAdmin = mock(CurrentAdmin.class);
        TemplateService service = new TemplateService(repository, fileAssetRepository, fileStorage, fileUrlService, currentAdmin);
        UUID fileId = UUID.randomUUID();
        when(fileAssetRepository.findById(fileId))
                .thenReturn(new FileAsset(fileId, "a.png", "INACTIVE", null));

        service.create(new TemplateService.TemplateCreateRequest("A", null, null, fileId));

        ArgumentCaptor<String> statusCap = ArgumentCaptor.forClass(String.class);
        ArgumentCaptor<UUID> fileIdCap = ArgumentCaptor.forClass(UUID.class);

        verify(repository).create(any(), any(), statusCap.capture(), any(), fileIdCap.capture(), isNull());
        verify(fileAssetRepository).updateStatus(fileId, "ACTIVE", null);

        assertEquals("ACTIVE", statusCap.getValue());
        assertEquals(fileId, fileIdCap.getValue());
    }

    @Test
    void deleteShouldRemoveTemplateImageAndFileRecord() throws Exception {
        TemplateRepository repository = mock(TemplateRepository.class);
        FileAssetRepository fileAssetRepository = mock(FileAssetRepository.class);
        FileStorage fileStorage = mock(FileStorage.class);
        FileUrlService fileUrlService = mock(FileUrlService.class);
        CurrentAdmin currentAdmin = mock(CurrentAdmin.class);
        TemplateService service = new TemplateService(repository, fileAssetRepository, fileStorage, fileUrlService, currentAdmin);
        UUID templateId = UUID.randomUUID();
        UUID fileId = UUID.randomUUID();
        UUID adminId = UUID.randomUUID();

        when(currentAdmin.idOrNull()).thenReturn(adminId);
        when(repository.findById(templateId))
                .thenReturn(new Template(templateId, "Template A", "ACTIVE", null, fileId, "a.png", null));
        when(fileAssetRepository.findById(fileId))
                .thenReturn(new FileAsset(fileId, "a.png", "ACTIVE", null));

        service.delete(templateId);

        verify(fileStorage).delete("a.png");
        verify(fileAssetRepository).deleteById(fileId);
        verify(repository).delete(templateId, adminId);
    }

    @Test
    void updateStatusShouldNormalizeToUppercase() {
        TemplateRepository repository = mock(TemplateRepository.class);
        FileAssetRepository fileAssetRepository = mock(FileAssetRepository.class);
        FileStorage fileStorage = mock(FileStorage.class);
        FileUrlService fileUrlService = mock(FileUrlService.class);
        CurrentAdmin currentAdmin = mock(CurrentAdmin.class);
        TemplateService service = new TemplateService(repository, fileAssetRepository, fileStorage, fileUrlService, currentAdmin);
        UUID templateId = UUID.randomUUID();
        UUID adminId = UUID.randomUUID();

        when(currentAdmin.idOrNull()).thenReturn(adminId);
        when(repository.findById(templateId))
                .thenReturn(new Template(templateId, "Template A", "ACTIVE", null, UUID.randomUUID(), "a.png", null));

        service.updateStatus(templateId, "inactive");

        verify(repository).updateStatus(templateId, "INACTIVE", adminId);
    }
}
