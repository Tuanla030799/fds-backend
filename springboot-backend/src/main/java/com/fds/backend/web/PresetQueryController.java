package com.fds.backend.web;

import com.fds.backend.domain.preset.PresetSummary;
import com.fds.backend.service.PresetQueryService;
import jakarta.validation.constraints.Max;
import jakarta.validation.constraints.Min;
import org.springframework.validation.annotation.Validated;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@Validated
@RestController
@RequestMapping("/api/v2/presets")
public class PresetQueryController {

    private final PresetQueryService presetQueryService;

    public PresetQueryController(PresetQueryService presetQueryService) {
        this.presetQueryService = presetQueryService;
    }

    @GetMapping
    public List<PresetSummary> list(
            @RequestParam(required = false) String status,
            @RequestParam(required = false) String keyword,
            @RequestParam(defaultValue = "1") @Min(1) int page,
            @RequestParam(defaultValue = "20") @Min(1) @Max(100) int size
    ) {
        return presetQueryService.search(status, keyword, page, size);
    }
}
