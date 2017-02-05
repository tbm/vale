Feature: Lint
  Scenario: Lint an AsciiDoc file
    When I lint "test.adoc"
    Then the output should contain exactly:
    """
    test.adoc:36:5:txtlint.Annotations:'NOTE' left in text
    test.adoc:43:1:txtlint.Annotations:'TODO' left in text
    test.adoc:47:4:txtlint.Annotations:'XXX' left in text

    """
    And the exit status should be 0

  Scenario: Lint a Python file
    When I lint "test.py"
    Then the output should contain exactly:
    """
    test.py:1:3:txtlint.Annotations:'FIXME' left in text
    test.py:5:5:txtlint.Annotations:'FIXME' left in text
    test.py:11:3:txtlint.Annotations:'XXX' left in text
    test.py:13:16:txtlint.Annotations:'XXX' left in text
    test.py:14:14:txtlint.Annotations:'NOTE' left in text
    test.py:17:1:txtlint.Annotations:'NOTE' left in text
    test.py:23:1:txtlint.Annotations:'XXX' left in text
    test.py:28:5:txtlint.Annotations:'NOTE' left in text
    test.py:35:8:txtlint.Annotations:'NOTE' left in text
    test.py:37:5:txtlint.Annotations:'TODO' left in text

    """
    And the exit status should be 0

  Scenario: Lint a C++ file
    When I lint "test.cc"
    Then the output should contain exactly:
    """
    test.cc:1:4:txtlint.Annotations:'XXX' left in text
    test.cc:9:6:txtlint.Annotations:'NOTE' left in text
    test.cc:13:6:txtlint.Annotations:'XXX' left in text
    test.cc:17:5:txtlint.Annotations:'FIXME' left in text
    test.cc:20:5:txtlint.Annotations:'XXX' left in text
    test.cc:23:37:txtlint.Annotations:'XXX' left in text

    """
    And the exit status should be 0

  Scenario: Lint a Markdown file
    When I lint "test.md"
    Then the output should contain exactly:
    """
    test.md:8:5:txtlint.Annotations:'NOTE' left in text
    test.md:18:4:txtlint.Annotations:'XXX' left in text

    """
    And the exit status should be 0

  Scenario: Lint a reStructuredText file
    When I lint "test.rst"
    Then the output should contain exactly:
    """
    test.rst:17:9:txtlint.Annotations:'NOTE' left in text
    test.rst:27:8:txtlint.Annotations:'XXX' left in text
    test.rst:50:29:txtlint.Annotations:'TODO' left in text
    test.rst:58:1:txtlint.Annotations:'NOTE' left in text

    """
    And the exit status should be 0

  Scenario: Lint a nonexistent file
    When I lint "null.cc"
    Then the output should contain exactly:
    """
    lstat null.cc: no such file or directory

    """
    And the exit status should be 1

  Scenario: Lint a Rust file
    When I lint "test.rs"
    Then the output should contain exactly:
    """
    test.rs:1:5:txtlint.Annotations:'NOTE' left in text
    test.rs:3:5:txtlint.Annotations:'XXX' left in text
    test.rs:5:17:txtlint.Annotations:'TODO' left in text
    test.rs:7:4:txtlint.Annotations:'FIXME' left in text
    test.rs:9:5:txtlint.Annotations:'XXX' left in text

    """
    And the exit status should be 0

  Scenario: Lint an R file
    When I lint "test.r"
    Then the output should contain exactly:
    """
    test.r:1:3:txtlint.Annotations:'NOTE' left in text
    test.r:6:22:txtlint.Annotations:'XXX' left in text

    """
    And the exit status should be 0

  Scenario: Lint a PHP file
    When I lint "test.php"
    Then the output should contain exactly:
    """
    test.php:2:31:txtlint.Annotations:'XXX' left in text
    test.php:3:8:txtlint.Annotations:'NOTE' left in text
    test.php:4:8:txtlint.Annotations:'FIXME' left in text
    test.php:6:33:txtlint.Annotations:'TODO' left in text

    """
    And the exit status should be 0

  Scenario: Lint a Lua file
    When I lint "test.lua"
    Then the output should contain exactly:
    """
    test.lua:1:4:txtlint.Annotations:'NOTE' left in text
    test.lua:2:19:txtlint.Annotations:'XXX' left in text
    test.lua:5:7:txtlint.Annotations:'NOTE' left in text
    test.lua:9:6:txtlint.Annotations:'XXX' left in text
    test.lua:15:4:txtlint.Annotations:'TODO' left in text

    """
    And the exit status should be 0

  Scenario: Lint a Haskell file
    When I lint "test.hs"
    Then the output should contain exactly:
    """
    test.hs:2:4:txtlint.Annotations:'NOTE' left in text
    test.hs:5:6:txtlint.Annotations:'TODO' left in text
    test.hs:6:25:txtlint.Annotations:'XXX' left in text
    test.hs:11:41:txtlint.Annotations:'XXX' left in text

    """
    And the exit status should be 0

  Scenario: Lint a Ruby file
    When I lint "test.rb"
    Then the output should contain exactly:
    """
    test.rb:2:1:txtlint.Annotations:'NOTE' left in text
    test.rb:6:1:txtlint.Annotations:'XXX' left in text
    test.rb:9:23:txtlint.Annotations:'XXX' left in text
    test.rb:11:3:txtlint.Annotations:'TODO' left in text

    """
    And the exit status should be 0
