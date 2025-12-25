# pdfx — a local-first PDF & document CLI 🔐

**pdfx** is a fast, privacy-friendly command-line tool for working with PDFs, images, and documents **entirely on your machine**.

No uploads.  
No tracking.  
No shady websites.  

Just one binary and predictable commands.

---

## Why pdfx?

Every time you fill a form or submit documents, you end up needing to:

- merge PDFs
- delete or reorder pages
- compress files under size limits
- convert images or documents to PDF

The usual solution?  
Uploading **sensitive documents** (IDs, contracts, resumes) to random websites.

**pdfx exists to fix this.**

> pdfx is a local alternative to tools like SmallPDF / iLovePDF —  
> but transparent, scriptable, and privacy-first.

---

## Key Principles

- 🔐 **Local-first** — all processing happens on your machine  
- 🚫 **No cloud, no telemetry, no tracking**
- ⚡ **Fast & scriptable** — designed for CLI workflows
- 📦 **Single Go binary** — easy to distribute
- 🧩 **Explicit dependencies** — no hidden magic

---

## Features

### 📄 PDF operations
- Merge multiple PDFs
- Delete pages (true delete, inclusive ranges)
- Split PDFs into pages
- Extract selected pages
- Compress PDFs (3 quality levels)

### 🖼️ Image operations
- Convert images → PDF
- Compress images (good / better / best)
- Convert image formats (jpg ↔ png)

### 📁 Document → PDF
- `.docx`, `.pptx`, `.xlsx` → PDF (LibreOffice)
- `.txt`, `.md`, `.html` → PDF (Pandoc)

### 🩺 System checks
- Built-in `doctor` command to verify dependencies

---

## Installation

First install the external dependency for sanity check use ```pdfx doctor``` then install the binary.

### External Dependency

Linux
```bash
sudo apt install ghostscript pandoc libreoffice
```

MacOS
```bash
brew install ghostscript pandoc
brew install --cask libreoffice
```

### Option 1 — Download binary (recommended)
Download the latest release from [GitHub Releases](https://github.com/yourusername/pdfx/releases) and place it in your `PATH`.

```bash
chmod +x pdfx
mv pdfx /usr/local/bin/
```

### Option 2 — Build from source

```bash
git clone https://github.com/yourusername/pdfx
cd pdfx
go build
```

---

## Quick Start

### Merge PDFs

```bash
pdfx merge a.pdf b.pdf c.pdf -o merged.pdf
```

### Delete pages (inclusive ranges)

```bash
pdfx delete report.pdf --pages 1,3-5,n -o cleaned.pdf
```

> Deletes page 1, pages 3-5, and the last page (`n`).

### Split a PDF

```bash
pdfx split report.pdf --pages 1-3
```

> Creates: `report_page_1.pdf`, `report_page_2.pdf`, `report_page_3.pdf`

### Compress a PDF

```bash
pdfx compress large.pdf --level best -o small.pdf
```

**Compression levels:**
- `good` — High quality, moderate compression
- `better` — Balanced (default)
- `best` — Smallest file size

### Images → PDF

```bash
pdfx img2pdf *.jpg -o images.pdf
```

**Supported formats:** `.jpg`, `.jpeg`, `.png`, `.webp`, `.bmp`, `.tif`, `.tiff`

### Convert documents to PDF

```bash
pdfx convert resume.docx -o resume.pdf
pdfx convert notes.md -o notes.pdf
```

### Compress images

```bash
pdfx img-compress photo.jpg --level best -o photo_small.jpg
```

### Convert image formats

```bash
pdfx img-convert image.webp -o image.png
```

---

## Page Range Syntax

Used by `delete`, `split`, and `extract` commands:

| Syntax | Meaning |
|--------|---------|
| `1` | Page 1 |
| `3-5` | Pages 3, 4, 5 (inclusive) |
| `n` | Last page |
| `2-n` | Page 2 through last page |
| `1,3-5,n` | Multiple ranges |

> ⚠️ **Note:** `1-n` (delete all pages) is not allowed in the `delete` command.

---

## Command Reference

### 🩺 `doctor` — System Dependency Check

```bash
pdfx doctor
```

Verifies that external tools are installed and available in your PATH.

---

### 📎 `merge` — Merge Multiple PDFs

```bash
pdfx merge <pdf1> <pdf2> [pdf3 ...] -o <output.pdf>
```

**Flags:**
- `-o`, `--output` — Output PDF file (default: `merged.pdf`)

---

### ❌ `delete` — Delete Pages from a PDF

```bash
pdfx delete <input.pdf> --pages <ranges> -o <output.pdf>
```

**Flags:**
- `--pages` — Pages to delete (required)
- `-o`, `--output` — Output PDF

**Examples:**

```bash
pdfx delete file.pdf --pages 3 -o out.pdf
pdfx delete file.pdf --pages 1,4-6,n -o cleaned.pdf
```

---

### ✂️ `split` — Extract Pages into Separate PDFs

```bash
pdfx split <input.pdf> --pages <ranges>
```

Extracts selected pages into individual PDF files.

**Example:**

```bash
pdfx split report.pdf --pages 1-3
```

---

### 📤 `extract` — Extract Pages into a Single PDF

```bash
pdfx extract <input.pdf> --pages <ranges> -o <output.pdf>
```

**Example:**

```bash
pdfx extract report.pdf --pages 2,4-6 -o excerpt.pdf
```

---

### 🖼️ `img2pdf` — Convert Images to PDF

```bash
pdfx img2pdf <image1> <image2> [...] -o <output.pdf>
```

**Examples:**

```bash
pdfx img2pdf a.jpg b.png -o images.pdf
pdfx img2pdf *.jpg -o album.pdf
```

---

### 🗜️ `compress` — Compress a PDF

```bash
pdfx compress <input.pdf> --level <good|better|best> -o <output.pdf>
```

**Flags:**
- `--level` — Compression level (default: `better`)
- `-o`, `--output` — Output PDF (default: `compressed.pdf`)

---

### 🔄 `convert` — Convert Documents to PDF

```bash
pdfx convert <input.file> -o <output.pdf>
```

**Supported formats:**
- Text: `.txt`, `.md`, `.html` (via Pandoc)
- Office: `.docx`, `.pptx`, `.xlsx`, `.odt` (via LibreOffice)

---

### 🖼️ `img-compress` — Compress Images

```bash
pdfx img-compress <image> --level <good|better|best> -o <output>
```

**Supported formats:**
- Input: `.jpg`, `.jpeg`, `.png`, `.webp`
- Output: `.jpg`, `.png`

---

### 🔁 `img-convert` — Convert Image Formats

```bash
pdfx img-convert <image> -o <output>
```

**Supported output:** `.jpg`, `.png`

---

## CLI Help

Every command has built-in help:

```bash
pdfx --help
pdfx delete --help
pdfx compress --help
```

---

## Security & Privacy

* ✅ No internet access
* ✅ No uploads
* ✅ No data collection
* ✅ No telemetry
* ✅ No background services

**You own your data. Always.**

---

## Roadmap

- [ ] Page reordering
- [ ] Batch operations
- [ ] Config file support
- [ ] Homebrew / Scoop packages
- [ ] Windows-friendly installers

---

## License

MIT License — use it, fork it, improve it.

---

## Final Note

This project was built to be **useful first**, clever second.

If it saved you time or protected your privacy —  
⭐ star the repo, or better, tell someone about it.

**Happy hacking 🚀**