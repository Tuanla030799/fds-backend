package com.fds.backend.preset;

import com.fds.backend.common.ApiResponse;
import jakarta.validation.Valid;
import jakarta.validation.constraints.Max;
import jakarta.validation.constraints.Min;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.UUID;

@RestController
@RequestMapping("/api")
public class PresetController {
    private final PresetService presetService;

    public PresetController(PresetService presetService) {
        this.presetService = presetService;
    }

    @GetMapping("/presets")
    public ApiResponse<List<Preset>> listPublic(@RequestParam(required = false) String status,
                                                @RequestParam(required = false) String keyword,
                                                @RequestParam(defaultValue = "1") @Min(1) int page,
                                                @RequestParam(defaultValue = "20") @Min(1) @Max(100) int limit) {
        return ApiResponse.ok("OK", presetService.list(status, keyword, page, limit));
    }

    @GetMapping("/admin/presets")
    public ApiResponse<List<Preset>> listAdmin(@RequestParam(required = false) String status,
                                               @RequestParam(required = false) String keyword,
                                               @RequestParam(defaultValue = "1") @Min(1) int page,
                                               @RequestParam(defaultValue = "20") @Min(1) @Max(100) int limit) {
        return ApiResponse.ok("OK", presetService.list(status, keyword, page, limit));
    }

    @PostMapping("/admin/presets")
    public ApiResponse<Void> create(@Valid @RequestBody PresetDtos.CreateRequest body) {
        presetService.create(new PresetService.PresetCreateRequest(
                body.name(),
                body.status(),
                body.note(),
                body.tags(),
                body.imageUrl(),
                body.sortOrder()
        ));
        return ApiResponse.ok("Created", null);
    }

    @DeleteMapping("/admin/presets/{id}")
    public ApiResponse<Void> delete(@PathVariable UUID id) {
        presetService.delete(id);
        return ApiResponse.ok("Deleted", null);
    }
}
