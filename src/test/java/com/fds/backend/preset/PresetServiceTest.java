package com.fds.backend.preset;

import org.junit.jupiter.api.Test;
import org.mockito.ArgumentCaptor;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.mockito.Mockito.*;

class PresetServiceTest {

    @Test
    void createShouldApplyDefaults() {
        PresetRepository repository = mock(PresetRepository.class);
        PresetService service = new PresetService(repository);

        service.create(new PresetService.PresetCreateRequest("A", null, null, null, "img", null));

        ArgumentCaptor<String> statusCap = ArgumentCaptor.forClass(String.class);
        ArgumentCaptor<String> tagsCap = ArgumentCaptor.forClass(String.class);
        ArgumentCaptor<Integer> sortCap = ArgumentCaptor.forClass(Integer.class);

        verify(repository).create(any(), any(), statusCap.capture(), any(), tagsCap.capture(), any(), sortCap.capture());

        assertEquals("active", statusCap.getValue());
        assertEquals("[]", tagsCap.getValue());
        assertEquals(0, sortCap.getValue());
    }
}
