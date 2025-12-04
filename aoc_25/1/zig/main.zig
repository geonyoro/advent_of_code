const std = @import("std");
const print = std.debug.print;
const expect = std.testing.expect;

pub fn main() !void {
    const file = try std.fs.cwd().openFile("input", .{});
    defer file.close();
    const pos = try file.getEndPos();
    const allocator = std.heap.page_allocator;
    const contents = try file.reader().readAllAlloc(allocator, pos);
    defer allocator.free(contents);
    var dial: i16 = 50;
    var it = std.mem.splitAny(u8, contents, "\n");
    var touchesZero: u16 = 0;
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const gallocator = gpa.allocator();
    while (it.next()) |line| {
        if (line.len == 0) {
            continue;
        }
        const moves = try std.fmt.parseInt(i16, line[1..], 10);
        var direction: i8 = 1;
        if (line[0] == 'L') {
            direction *= -1;
        }
        const val = try moveDial(gallocator, dial, moves, direction);
        defer gallocator.free(val.debugString);
        dial = val.dial;
        touchesZero += val.touchesZero;
    }
    print("Zero: {d}\n", .{touchesZero});
}

const DialMove = struct { dial: i16, touchesZero: u16, debugString: []u8 };
fn moveDial(allocator: std.mem.Allocator, dial_: i16, moves_: i16, direction: i8) !DialMove {
    var dial = dial_;
    var moves = moves_;
    var touchesZero: u16 = 0;
    const rotations: u16 = @divTrunc(@abs(moves), 100);
    moves = @mod(moves, 100);
    moves *= direction;
    touchesZero += rotations;
    dial = dial + moves;
    if (dial < 0) {
        dial = 100 + dial;
        if (dial_ > 0) {
            // this one didn't start by touching zero
            touchesZero += 1;
        }
    } else if (dial >= 100) {
        dial = dial - 100;
        touchesZero += 1;
    } else if (dial == 0) {
        touchesZero += 1;
    }
    // for later printing if needed
    const string = try std.fmt.allocPrint(allocator, "dial:{d} rot:{d} change:{d}", .{ dial, rotations, moves });
    return .{ .dial = dial, .touchesZero = touchesZero, .debugString = string };
}

test "moveDial" {
    const testCase = struct { startDial: i16, moves: i16, dial: i16, touchesZero: i16 };
    const cases = [_]testCase{
        testCase{ .startDial = 50, .moves = -1000, .dial = 50, .touchesZero = 10 },
        testCase{ .startDial = 50, .moves = 1000, .dial = 50, .touchesZero = 10 },
        testCase{ .startDial = 50, .moves = -68, .dial = 82, .touchesZero = 1 },
        testCase{ .startDial = 82, .moves = -30, .dial = 52, .touchesZero = 0 },
        testCase{ .startDial = 52, .moves = 48, .dial = 0, .touchesZero = 1 },
        testCase{ .startDial = 0, .moves = -5, .dial = 95, .touchesZero = 0 },
        testCase{ .startDial = 95, .moves = 60, .dial = 55, .touchesZero = 1 },
        testCase{ .startDial = 55, .moves = -55, .dial = 0, .touchesZero = 1 },
        testCase{ .startDial = 0, .moves = -1, .dial = 99, .touchesZero = 0 },
        testCase{ .startDial = 99, .moves = -99, .dial = 0, .touchesZero = 1 },
        testCase{ .startDial = 0, .moves = 14, .dial = 14, .touchesZero = 0 },
        testCase{ .startDial = 14, .moves = -82, .dial = 32, .touchesZero = 1 },
    };
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const allocator = gpa.allocator();
    for (cases) |c| {
        var temp_moves: i16 = c.moves;
        if (c.moves < 0) {
            temp_moves = 0 - temp_moves;
        }
        var dir: i8 = 1;
        if (c.moves < 0) {
            dir = -1;
        }
        const val = try moveDial(allocator, c.startDial, temp_moves, dir);
        defer allocator.free(val.debugString);
        errdefer {
            print("{any}\n", .{c});
            print("{s}\n", .{val.debugString});
        }
        try expect(c.touchesZero == val.touchesZero);
        try expect(c.dial == val.dial);
    }
}
