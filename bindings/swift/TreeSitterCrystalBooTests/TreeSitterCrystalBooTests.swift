import XCTest
import SwiftTreeSitter
import TreeSitterCrystalBoo

final class TreeSitterCrystalBooTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_crystal_boo())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading CrystalBoo grammar")
    }
}
