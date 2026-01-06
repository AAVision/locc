package main

// Language represents a programming language with its comment patterns
type Language struct {
	Name              string
	Extensions        []string
	SingleLineComment string
	MultiLineStart    string
	MultiLineEnd      string
}

// Languages defines all supported programming languages and their comment patterns
var Languages = map[string]*Language{
	".go": {
		Name:              "Go",
		Extensions:        []string{".go"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	},
	".js": {
		Name:              "JavaScript",
		Extensions:        []string{".js"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	},
	".ts": {
		Name:              "TypeScript",
		Extensions:        []string{".ts"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	},
	".tsx": {
		Name:              "TypeScript JSX",
		Extensions:        []string{".tsx"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	},
	".jsx": {
		Name:              "JavaScript JSX",
		Extensions:        []string{".jsx"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	},
	".html": {
		Name:              "HTML",
		Extensions:        []string{".html", ".htm"},
		SingleLineComment: "",
		MultiLineStart:    "<!--",
		MultiLineEnd:      "-->",
	},
	".htm": {
		Name:              "HTML",
		Extensions:        []string{".html", ".htm"},
		SingleLineComment: "",
		MultiLineStart:    "<!--",
		MultiLineEnd:      "-->",
	},
	".py": {
		Name:              "Python",
		Extensions:        []string{".py"},
		SingleLineComment: "#",
		MultiLineStart:    `"""`,
		MultiLineEnd:      `"""`,
	},
	".rb": {
		Name:              "Ruby",
		Extensions:        []string{".rb"},
		SingleLineComment: "#",
		MultiLineStart:    "=begin",
		MultiLineEnd:      "=end",
	},
	".java": {
		Name:              "Java",
		Extensions:        []string{".java"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	},
	".c": {
		Name:              "C",
		Extensions:        []string{".c"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	},
	".h": {
		Name:              "C Header",
		Extensions:        []string{".h"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	},
	".cpp": {
		Name:              "C++",
		Extensions:        []string{".cpp", ".cc", ".cxx"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	},
	".cc": {
		Name:              "C++",
		Extensions:        []string{".cpp", ".cc", ".cxx"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	},
	".hpp": {
		Name:              "C++ Header",
		Extensions:        []string{".hpp"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	},
	".cs": {
		Name:              "C#",
		Extensions:        []string{".cs"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	},
	".php": {
		Name:              "PHP",
		Extensions:        []string{".php"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	},
	".swift": {
		Name:              "Swift",
		Extensions:        []string{".swift"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	},
	".kt": {
		Name:              "Kotlin",
		Extensions:        []string{".kt"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	},
	".rs": {
		Name:              "Rust",
		Extensions:        []string{".rs"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	},
	".scala": {
		Name:              "Scala",
		Extensions:        []string{".scala"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	},
	".json": {
		Name:              "JSON",
		Extensions:        []string{".json"},
		SingleLineComment: "",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	".yaml": {
		Name:              "YAML",
		Extensions:        []string{".yaml", ".yml"},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	".yml": {
		Name:              "YAML",
		Extensions:        []string{".yaml", ".yml"},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	".md": {
		Name:              "Markdown",
		Extensions:        []string{".md"},
		SingleLineComment: "",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	".css": {
		Name:              "CSS",
		Extensions:        []string{".css"},
		SingleLineComment: "",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	},
	".scss": {
		Name:              "SCSS",
		Extensions:        []string{".scss"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	},
	".sass": {
		Name:              "Sass",
		Extensions:        []string{".sass"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	},
	".less": {
		Name:              "Less",
		Extensions:        []string{".less"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	},
	".sql": {
		Name:              "SQL",
		Extensions:        []string{".sql"},
		SingleLineComment: "--",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	},
	".sh": {
		Name:              "Shell",
		Extensions:        []string{".sh", ".bash"},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	".bash": {
		Name:              "Shell",
		Extensions:        []string{".sh", ".bash"},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	".xml": {
		Name:              "XML",
		Extensions:        []string{".xml"},
		SingleLineComment: "",
		MultiLineStart:    "<!--",
		MultiLineEnd:      "-->",
	},
	".vue": {
		Name:              "Vue",
		Extensions:        []string{".vue"},
		SingleLineComment: "//",
		MultiLineStart:    "<!--",
		MultiLineEnd:      "-->",
	},
	".svelte": {
		Name:              "Svelte",
		Extensions:        []string{".svelte"},
		SingleLineComment: "//",
		MultiLineStart:    "<!--",
		MultiLineEnd:      "-->",
	},
	".lua": {
		Name:              "Lua",
		Extensions:        []string{".lua"},
		SingleLineComment: "--",
		MultiLineStart:    "--[[",
		MultiLineEnd:      "]]",
	},
	".r": {
		Name:              "R",
		Extensions:        []string{".r", ".R"},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	".R": {
		Name:              "R",
		Extensions:        []string{".r", ".R"},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	".pl": {
		Name:              "Perl",
		Extensions:        []string{".pl", ".pm"},
		SingleLineComment: "#",
		MultiLineStart:    "=pod",
		MultiLineEnd:      "=cut",
	},
	".pm": {
		Name:              "Perl",
		Extensions:        []string{".pl", ".pm"},
		SingleLineComment: "#",
		MultiLineStart:    "=pod",
		MultiLineEnd:      "=cut",
	},
	".ex": {
		Name:              "Elixir",
		Extensions:        []string{".ex", ".exs"},
		SingleLineComment: "#",
		MultiLineStart:    `"""`,
		MultiLineEnd:      `"""`,
	},
	".exs": {
		Name:              "Elixir",
		Extensions:        []string{".ex", ".exs"},
		SingleLineComment: "#",
		MultiLineStart:    `"""`,
		MultiLineEnd:      `"""`,
	},
	".erl": {
		Name:              "Erlang",
		Extensions:        []string{".erl"},
		SingleLineComment: "%",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	".hs": {
		Name:              "Haskell",
		Extensions:        []string{".hs"},
		SingleLineComment: "--",
		MultiLineStart:    "{-",
		MultiLineEnd:      "-}",
	},
	".clj": {
		Name:              "Clojure",
		Extensions:        []string{".clj"},
		SingleLineComment: ";",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	".toml": {
		Name:              "TOML",
		Extensions:        []string{".toml"},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	".ini": {
		Name:              "INI",
		Extensions:        []string{".ini"},
		SingleLineComment: ";",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	".dockerfile": {
		Name:              "Dockerfile",
		Extensions:        []string{".dockerfile"},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	".makefile": {
		Name:              "Makefile",
		Extensions:        []string{".makefile"},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	".tf": {
		Name:              "Terraform",
		Extensions:        []string{".tf"},
		SingleLineComment: "#",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	},
	".proto": {
		Name:              "Protocol Buffers",
		Extensions:        []string{".proto"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	},
	".graphql": {
		Name:              "GraphQL",
		Extensions:        []string{".graphql", ".gql"},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	".gql": {
		Name:              "GraphQL",
		Extensions:        []string{".graphql", ".gql"},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	".txt": {
		Name:              "Text",
		Extensions:        []string{".txt"},
		SingleLineComment: "",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	".hcl": {
		Name:              "HCL",
		Extensions:        []string{".hcl"},
		SingleLineComment: "#",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	},
	".y": {
		Name:              "Yacc",
		Extensions:        []string{".y"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	},
	".nix": {
		Name:              "Nix",
		Extensions:        []string{".nix"},
		SingleLineComment: "#",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	},
	".json5": {
		Name:              "JSON5",
		Extensions:        []string{".json5"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	},
}

// BinaryExtensions contains file extensions that should be skipped
var BinaryExtensions = map[string]bool{
	// Images
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".gif":  true,
	".bmp":  true,
	".ico":  true,
	".svg":  true,
	".webp": true,
	".tiff": true,
	".tif":  true,
	".psd":  true,
	".raw":  true,
	".heif": true,
	".heic": true,
	".avif": true,

	// Documents
	".pdf":  true,
	".doc":  true,
	".docx": true,
	".xls":  true,
	".xlsx": true,
	".ppt":  true,
	".pptx": true,
	".odt":  true,
	".ods":  true,
	".odp":  true,
	".rtf":  true,
	".epub": true,
	".mobi": true,

	// Archives
	".zip":  true,
	".tar":  true,
	".gz":   true,
	".tgz":  true,
	".bz2":  true,
	".xz":   true,
	".lz":   true,
	".lzma": true,
	".rar":  true,
	".7z":   true,
	".cab":  true,
	".iso":  true,
	".dmg":  true,
	".deb":  true,
	".rpm":  true,
	".apk":  true,
	".msi":  true,

	// Executables and libraries
	".exe":   true,
	".dll":   true,
	".so":    true,
	".dylib": true,
	".bin":   true,
	".com":   true,
	".mach":  true,
	".elf":   true,

	// Data files
	".dat":     true,
	".db":      true,
	".sqlite":  true,
	".sqlite3": true,
	".mdb":     true,
	".accdb":   true,
	".frm":     true,
	".ibd":     true,
	".dbf":     true,

	// Audio
	".mp3":  true,
	".wav":  true,
	".flac": true,
	".ogg":  true,
	".aac":  true,
	".wma":  true,
	".m4a":  true,
	".aiff": true,
	".mid":  true,
	".midi": true,

	// Video
	".mp4":  true,
	".avi":  true,
	".mov":  true,
	".wmv":  true,
	".flv":  true,
	".mkv":  true,
	".webm": true,
	".m4v":  true,
	".mpeg": true,
	".mpg":  true,
	".3gp":  true,

	// Fonts
	".ttf":   true,
	".otf":   true,
	".woff":  true,
	".woff2": true,
	".eot":   true,
	".fon":   true,

	// Compiled/bytecode
	".class": true,
	".jar":   true,
	".war":   true,
	".ear":   true,
	".pyc":   true,
	".pyo":   true,
	".pyd":   true,
	".o":     true,
	".a":     true,
	".obj":   true,
	".lib":   true,
	".ko":    true,
	".beam":  true,
	".elc":   true,

	// Lock files (often auto-generated)
	".lock": true,

	// Node.js specific
	".node": true,

	// Other binary formats
	".wasm":    true,
	".min.js":  true,
	".min.css": true,
	".map":     true,
	".pak":     true,
	".cache":   true,
	".swp":     true,
	".swo":     true,
}

// FilenameLanguages maps specific filenames (without extension) to languages
var FilenameLanguages = map[string]*Language{
	"Makefile": {
		Name:              "Makefile",
		Extensions:        []string{},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	"makefile": {
		Name:              "Makefile",
		Extensions:        []string{},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	"GNUmakefile": {
		Name:              "Makefile",
		Extensions:        []string{},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	"Dockerfile": {
		Name:              "Dockerfile",
		Extensions:        []string{},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	"dockerfile": {
		Name:              "Dockerfile",
		Extensions:        []string{},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	"LICENSE": {
		Name:              "License",
		Extensions:        []string{},
		SingleLineComment: "",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	"LICENSE.txt": {
		Name:              "License",
		Extensions:        []string{},
		SingleLineComment: "",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	"LICENSE.md": {
		Name:              "License",
		Extensions:        []string{},
		SingleLineComment: "",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	"LICENCE": {
		Name:              "License",
		Extensions:        []string{},
		SingleLineComment: "",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	"COPYING": {
		Name:              "License",
		Extensions:        []string{},
		SingleLineComment: "",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	"README": {
		Name:              "Readme",
		Extensions:        []string{},
		SingleLineComment: "",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	"README.txt": {
		Name:              "Readme",
		Extensions:        []string{},
		SingleLineComment: "",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	"Vagrantfile": {
		Name:              "Vagrantfile",
		Extensions:        []string{},
		SingleLineComment: "#",
		MultiLineStart:    "=begin",
		MultiLineEnd:      "=end",
	},
	"Gemfile": {
		Name:              "Gemfile",
		Extensions:        []string{},
		SingleLineComment: "#",
		MultiLineStart:    "=begin",
		MultiLineEnd:      "=end",
	},
	"Rakefile": {
		Name:              "Rakefile",
		Extensions:        []string{},
		SingleLineComment: "#",
		MultiLineStart:    "=begin",
		MultiLineEnd:      "=end",
	},
	"Procfile": {
		Name:              "Procfile",
		Extensions:        []string{},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	"CMakeLists.txt": {
		Name:              "CMake",
		Extensions:        []string{},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	"Jenkinsfile": {
		Name:              "Jenkinsfile",
		Extensions:        []string{},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	},
	"CHANGELOG": {
		Name:              "Changelog",
		Extensions:        []string{},
		SingleLineComment: "",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	"CHANGELOG.md": {
		Name:              "Changelog",
		Extensions:        []string{},
		SingleLineComment: "",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	"AUTHORS": {
		Name:              "Authors",
		Extensions:        []string{},
		SingleLineComment: "",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	"CONTRIBUTORS": {
		Name:              "Contributors",
		Extensions:        []string{},
		SingleLineComment: "",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
}

// HiddenFileLanguages maps hidden config files to languages
var HiddenFileLanguages = map[string]*Language{
	".gitignore": {
		Name:              "Git Config",
		Extensions:        []string{},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	".gitattributes": {
		Name:              "Git Config",
		Extensions:        []string{},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	".dockerignore": {
		Name:              "Docker Config",
		Extensions:        []string{},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	".editorconfig": {
		Name:              "EditorConfig",
		Extensions:        []string{},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	".eslintrc": {
		Name:              "ESLint Config",
		Extensions:        []string{},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	},
	".prettierrc": {
		Name:              "Prettier Config",
		Extensions:        []string{},
		SingleLineComment: "",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	".babelrc": {
		Name:              "Babel Config",
		Extensions:        []string{},
		SingleLineComment: "",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	".npmrc": {
		Name:              "NPM Config",
		Extensions:        []string{},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	".yarnrc": {
		Name:              "Yarn Config",
		Extensions:        []string{},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	".env": {
		Name:              "Environment",
		Extensions:        []string{},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	".env.example": {
		Name:              "Environment",
		Extensions:        []string{},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	".env.local": {
		Name:              "Environment",
		Extensions:        []string{},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	".htaccess": {
		Name:              "Apache Config",
		Extensions:        []string{},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
	".travis.yml": {
		Name:              "Travis CI",
		Extensions:        []string{},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	},
}

// GetLanguage returns the language definition for a given file extension
func GetLanguage(ext string) *Language {
	if lang, ok := Languages[ext]; ok {
		return lang
	}
	return nil
}

// GetLanguageByFilename returns the language definition for a specific filename
func GetLanguageByFilename(filename string) *Language {
	// Check exact filename match first
	if lang, ok := FilenameLanguages[filename]; ok {
		return lang
	}
	// Check hidden file languages
	if lang, ok := HiddenFileLanguages[filename]; ok {
		return lang
	}
	return nil
}

// IsBinaryExtension checks if the file extension is a binary file
func IsBinaryExtension(ext string) bool {
	return BinaryExtensions[ext]
}
