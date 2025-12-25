# 📘 pdfx Command Line Reference

**pdfx** is a local-first, privacy-friendly CLI for working with PDFs and images.  
All operations happen on your machine — **no uploads, no tracking**.

---

## 🚀 Getting Started

### General Usage

```bash
pdfx <command> [arguments] [flags]
```

### Get Help

```bash
pdfx --help
pdfx <command> --help
```

---

## 📋 Commands

### 🩺 `doctor` — System Dependency Check

Checks whether required external tools are installed and available in PATH.

```bash
pdfx doctor
```

#### Dependencies

| Dependency | Used for |
|------------|----------|
| **Ghostscript** | PDF compression |
| **Pandoc** | `.txt`, `.md`, `.html` → PDF |
| **LibreOffice** | `.docx`, `.pptx`, `.xlsx` → PDF |

#### Example Output

```
pdfx system check:
  ✔ Ghostscript (PDF compression)
  ✔ Pandoc (Text/Markdown to PDF)
  ✘ LibreOffice (DOCX/PPTX/XLSX to PDF) — NOT FOUND
```

---

### 📎 `merge` — Merge Multiple PDFs

Merge two or more PDF files into a single PDF.

```bash
pdfx merge <pdf1> <pdf2> [pdf3 ...] -o <output.pdf>
```

#### Flags

| Flag | Description | Default |
|------|-------------|---------|
| `-o`, `--output` | Output PDF file | `merged.pdf` |

#### Example

```bash
pdfx merge a.pdf b.pdf c.pdf -o combined.pdf
```

---

### ❌ `delete` — Delete Pages from a PDF

Deletes pages from a PDF using inclusive page ranges.

```bash
pdfx delete <input.pdf> --pages <ranges> -o <output.pdf>
```

#### Page Range Syntax

| Syntax | Meaning |
|--------|---------|
| `1` | Page 1 |
| `3-5` | Pages 3, 4, 5 (inclusive) |
| `n` | Last page |
| `2-n` | Page 2 through last page |
| `1,3-5,n` | Multiple ranges |

> ⚠️ **Note:** `1-n` (delete all pages) is not allowed.

#### Flags

| Flag | Description |
|------|-------------|
| `--pages` | Pages to delete (required) |
| `-o`, `--output` | Output PDF |

#### Examples

```bash
pdfx delete file.pdf --pages 3 -o out.pdf
pdfx delete file.pdf --pages 1,4-6,n -o cleaned.pdf
```

---

### ✂️ `split` — Extract Pages into Separate PDFs

Extract selected pages into individual PDF files.

```bash
pdfx split <input.pdf> --pages <ranges>
```

#### Page Range Syntax

Same as `delete`, but interpreted as pages to **KEEP**.

#### Output Behavior

Creates files like:
- `input_page_1.pdf`
- `input_page_2.pdf`

#### Example

```bash
pdfx split report.pdf --pages 1-3
```

---

### 📤 `extract` — Extract Pages into a Single PDF

Extract selected pages and write them into one new PDF.

```bash
pdfx extract <input.pdf> --pages <ranges> -o <output.pdf>
```

#### Example

```bash
pdfx extract report.pdf --pages 2,4-6 -o excerpt.pdf
```

---

### 🖼️ `img2pdf` — Convert Images to PDF

Convert one or more images into a single PDF.

```bash
pdfx img2pdf <image1> <image2> [...] -o <output.pdf>
```

#### Supported Input Formats

- `.jpg`, `.jpeg`
- `.png`
- `.webp`
- `.bmp`
- `.tif`, `.tiff`

#### Notes

- Image order is preserved as provided
- Mixed formats are allowed

#### Examples

```bash
pdfx img2pdf a.jpg b.png -o images.pdf
pdfx img2pdf *.jpg -o album.pdf
```

---

### 🗜️ `compress` — Compress a PDF

Compress a PDF using Ghostscript with preset quality levels.

```bash
pdfx compress <input.pdf> --level <good|better|best> -o <output.pdf>
```

#### Compression Levels

| Level | Description |
|-------|-------------|
| `good` | High quality, moderate compression |
| `better` | Balanced (default) |
| `best` | Smallest file size |

#### Flags

| Flag | Description | Default |
|------|-------------|---------|
| `--level` | Compression level | `better` |
| `-o`, `--output` | Output PDF | `compressed.pdf` |

#### Examples

```bash
pdfx compress file.pdf --level good
pdfx compress file.pdf --level best -o small.pdf
```

---

### 🔄 `convert` — Convert Documents to PDF

Convert common document formats to PDF.

```bash
pdfx convert <input.file> -o <output.pdf>
```

#### Supported Formats

| Type | Formats | Backend |
|------|---------|---------|
| **Text** | `.txt`, `.md`, `.html` | Pandoc |
| **Office** | `.docx`, `.pptx`, `.xlsx`, `.odt` | LibreOffice |

#### Examples

```bash
pdfx convert resume.docx -o resume.pdf
pdfx convert notes.md -o notes.pdf
```

---

### 🖼️ `img-compress` — Compress Images

Compress images locally using quality presets.

```bash
pdfx img-compress <image> --level <good|better|best> -o <output>
```

#### Supported Formats

- **Input:** `.jpg`, `.jpeg`, `.png`, `.webp`
- **Output:** `.jpg`, `.png`

#### Examples

```bash
pdfx img-compress photo.jpg --level best -o small.jpg
pdfx img-compress diagram.png --level good -o diagram.png
```

---

### 🔁 `img-convert` — Convert Image Formats

Convert images between supported formats.

```bash
pdfx img-convert <image> -o <output>
```

#### Supported Output Formats

- `.jpg`, `.png`

#### Examples

```bash
pdfx img-convert photo.webp -o photo.png
pdfx img-convert scan.jpg -o scan.png
```

---

## 🆘 Help

Built-in help is available for every command:

```bash
pdfx help
pdfx help delete
pdfx delete --help
```

---

## 🔒 Privacy First

All operations run locally on your machine. **No data is uploaded to external servers.**

---

## 📄 License

[Add your license information here]

## 🤝 Contributing

[Add contribution guidelines here]

---

**Made with ❤️ for privacy-conscious PDF users**