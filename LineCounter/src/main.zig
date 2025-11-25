const std = @import("std");
const fs = std.fs;
const print = std.debug.print;

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer if (gpa.deinit() != .ok) @panic("leak");
    const allocator = gpa.allocator();
    const ignoreDirs = [_][]const u8{ ".git", "zig-out", "node_modules" };

    // var count = 0;

    // In order to walk the directry, `iterate` must be set to true.
    var dir = try fs.cwd().openDir("../../", .{ .iterate = true });
    defer dir.close();

    var walker = try dir.walk(allocator);
    defer walker.deinit();

    while (try walker.next()) |entry| {
        if(entry.kind == .Dir) {
            var should_ignore = false;
            for (ignoreDirs) |ignoreDir| {
                if (entry.basename == ignoreDir) {
                    should_ignore = true;
                    break;
                }
            }
            if (should_ignore) {
                walker.skip();
                continue;
            }
        }
        print("path: {s}, basename:{s}, type:{s}\n", .{
            entry.path,
            entry.basename,
            @tagName(entry.kind),
        });
    }
}
