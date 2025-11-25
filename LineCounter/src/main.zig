const std = @import("std");
const fs = std.fs;
const print = std.debug.print;

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer if (gpa.deinit() != .ok) @panic("leak");
    const allocator = gpa.allocator();
    const ignoreDirs = [_][]const u8{ ".git", "zig-out", "node_modules" };
    const extensions = [_][]const u8{ ".go", ".vint", ".h", ".hpp", ".S" };

    var total_lines: u64 = 0;

    // Open the current directory (change to "." for current dir)
    var dir = try fs.cwd().openDir(".", .{ .iterate = true });
    defer dir.close();

    var walker = try dir.walk(allocator);
    defer walker.deinit();

    while (try walker.next()) |entry| {
        if (entry.kind == .directory) {
            var should_ignore = false;
            for (ignoreDirs) |ignoreDir| {
                if (std.mem.eql(u8, entry.basename, ignoreDir)) {
                    should_ignore = true;
                    break;
                }
            }
            if (should_ignore) {
                try walker.skip();
                continue;
            }
        } else if (entry.kind == .file) {
            // Check if file matches any extension
            var matches = false;
            for (extensions) |ext| {
                if (std.mem.endsWith(u8, entry.basename, ext)) {
                    matches = true;
                    break;
                }
            }
            if (!matches) continue;

            // Open the file and count lines
            const file = try dir.dir.openFile(entry.path, .{});
            defer file.close();

            var buf_reader = std.io.bufferedReader(file.reader());
            const reader = buf_reader.reader();

            var line_count: u64 = 0;
            while (true) {
                var buf: [4096]u8 = undefined;
                const bytes_read = try reader.read(&buf);
                if (bytes_read == 0) break;

                for (buf[0..bytes_read]) |byte| {
                    if (byte == '\n') line_count += 1;
                }
            }

            // If file doesn't end with \n, count the last line
            const stat = try file.stat();
            if (stat.size > 0) {
                line_count += 1; // Assume files end with content even without trailing \n
            }

            total_lines += line_count;

            print("File: {s}, lines: {}\n", .{ entry.path, line_count });
        }
    }

    print("Total lines: {}\n", .{total_lines});
}
