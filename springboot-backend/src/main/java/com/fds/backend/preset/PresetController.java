package com.fds.backend.preset;

import com.fds.backend.common.ApiResponse;
import jakarta.validation.constraints.Max;
import jakarta.validation.constraints.Min;
import jakarta.validation.constraints.NotBlank;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.UUID;

@RestController
public class PresetController {
    private final PresetService presetService;

    public PresetController(PresetService presetService) {
        this.presetService = presetService;
    }

    @GetMapping("/api/presets")
    public ApiResponse<List<Preset>> listPublic(@RequestParam(required = false) String status,
                                                @RequestParam(required = false) String keyword,
                                                @RequestParam(defaultValue = "1") @Min(1) int page,
                                                @RequestParam(defaultValue = "20") @Min(1) @Max(100) int limit) {
        return ApiResponse.ok("OK", presetService.list(status, keyword, page, limit));
    }

    @GetMapping("/api/admin/presets")
    public ApiResponse<List<Preset>> listAdmin(@RequestParam(required = false) String status,
                                               @RequestParam(required = false) String keyword,
                                               @RequestParam(defaultValue = "1") @Min(1) int page,
                                               @RequestParam(defaultValue = "20") @Min(1) @Max(100) int limit) {
        return ApiResponse.ok("OK", presetService.list(status, keyword, page, limit));
    }

    @PostMapping("/api/admin/presets")
    public ApiResponse<Void> create(@RequestBody CreatePresetBody body) {
        presetService.create(new PresetService.PresetCreateRequest(body.name, body.status, body.note,
                body.tags == null ? "[]" : body.tags, body.imageUrl, body.sortOrder));
        return ApiResponse.ok("Created", null);
    }

    @DeleteMapping("/api/admin/presets/{id}")
    public ApiResponse<Void> delete(@PathVariable UUID id) {
        presetService.delete(id);
        return ApiResponse.ok("Deleted", null);
    }

    public static class CreatePresetBody {
        @NotBlank public String name;
        public String status = "active";
        public String note;
        public String tags;
        @NotBlank public String imageUrl;
        public int sortOrder = 0;
    }
}
