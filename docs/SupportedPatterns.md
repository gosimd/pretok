# Supported patterns

This is the documentation registry of all **23 supported exact regex strings**,
ordered by corpus ID. The [runtime registry](../internal/compile/registry.go)
selects these expressions; the [25-string corpus](../testdata/corpus_patterns.json)
also contains unsupported targets P14 and P25. IDs are stable.

`regexp(r"...")` is documentation notation for the string passed to `pretok.Compile`,
not an additional public API. Raw string literals preserve every backslash.
For P09, P11 and P22, normal `"\r\n"` literals are concatenated with raw fragments
to represent actual CR/LF characters without hiding them in the Markdown source.
These expressions use Python string-literal syntax; no regex parsing or
normalization is implied. In particular, P05 and P09 are distinct exact strings.

Each entry lists at most three tokenizer examples from the recorded research
snapshot, with links to immutable file revisions and the location of the regex.
Examples are not a popularity ranking or a claim about the current `main` branch.
A tokenizer can use several listed patterns in a Sequence. Recognizing its regex
does not imply support for its entire tokenizer pipeline; see [Regexp](Regexp.md)
for matching, gap handling and compatibility contracts.

Jump to: [P01](#p01) · [P02](#p02) · [P03](#p03) · [P04](#p04) · [P05](#p05) · [P06](#p06) · [P07](#p07) · [P08](#p08) · [P09](#p09) · [P10](#p10) · [P11](#p11) · [P12](#p12) · [P13](#p13) · [P15](#p15) · [P16](#p16) · [P17](#p17) · [P18](#p18) · [P19](#p19) · [P20](#p20) · [P21](#p21) · [P22](#p22) · [P23](#p23) · [P24](#p24).

## P01

```python
regexp(r"(?i:'s|'t|'re|'ve|'m|'ll|'d)|[^\r\n\p{L}\p{N}]?[\p{L}\p{M}]+|\p{N}| ?[^\s\p{L}\p{M}\p{N}]+[\r\n]*|\s*[\r\n]+|\s+(?!\S)|\s+")
```

[Exact string and validation provenance](../testdata/patterns/P01/manifest.json).

Tokenizer examples:

- [Qwen/Qwen3.5-0.8B](https://huggingface.co/Qwen/Qwen3.5-0.8B/blob/2fc06364715b967f1860aea9cf38778875588b17/tokenizer.json) — `/pre_tokenizer/pretokenizers/0/pattern/Regex`.
- [Qwen/Qwen3.6-27B](https://huggingface.co/Qwen/Qwen3.6-27B/blob/6a9e13bd6fc8f0983b9b99948120bc37f49c13e9/tokenizer.json) — `/pre_tokenizer/pretokenizers/0/pattern/Regex`.
- [Qwen/Qwen3.8-27B](https://huggingface.co/Qwen/Qwen3.8-27B/blob/1d4bf0f2ff6012fd82039f2fa52739d0dd7c60c0/tokenizer.json) — `/pre_tokenizer/pretokenizers/0/pattern/Regex`.

## P02

```python
regexp(r"[0-9][0-9][0-9]")
```

[Exact string and validation provenance](../testdata/patterns/P02/manifest.json).

Tokenizer examples:

- [tiiuae/falcon-7b](https://huggingface.co/tiiuae/falcon-7b/blob/ec89142b67d748a1865ea4451372db8313ada0d8/tokenizer.json) — `/pre_tokenizer/pretokenizers/3/pattern/Regex`.
- [tiiuae/falcon-7b-instruct](https://huggingface.co/tiiuae/falcon-7b-instruct/blob/8782b5c5d8c9290412416618f36a133653e85285/tokenizer.json) — `/pre_tokenizer/pretokenizers/3/pattern/Regex`.
- [tiiuae/falcon-mamba-7b](https://huggingface.co/tiiuae/falcon-mamba-7b/blob/080ad94b3619e2c2d0afa59bafdc6113465b7006/tokenizer.json) — `/pre_tokenizer/pretokenizers/3/pattern/Regex`.

## P03

```python
regexp(r"\s+$")
```

[Exact string and validation provenance](../testdata/patterns/P03/manifest.json).

Tokenizer examples:

- [deepseek-ai/DeepSeek-Coder-V2-Lite-Instruct](https://huggingface.co/deepseek-ai/DeepSeek-Coder-V2-Lite-Instruct/blob/e434a23f91ba5b4923cf6c9d9a238eb4a08e3a11/tokenizer.json) — `/pre_tokenizer/pretokenizers/3/pattern/Regex`.
- [deepseek-ai/DeepSeek-V2](https://huggingface.co/deepseek-ai/DeepSeek-V2/blob/4461458f186c35188585855f28f77af5661ad489/tokenizer.json) — `/pre_tokenizer/pretokenizers/3/pattern/Regex`.
- [deepseek-ai/deepseek-coder-7b-instruct-v1.5](https://huggingface.co/deepseek-ai/deepseek-coder-7b-instruct-v1.5/blob/2a050a4c59d687a85324d32e147517992117ed30/tokenizer.json) — `/pre_tokenizer/pretokenizers/3/pattern/Regex`.

## P04

```python
regexp(r"[^\r\n\p{L}\p{N}]?[\p{Lu}\p{Lt}\p{Lm}\p{Lo}\p{M}]*[\p{Ll}\p{Lm}\p{Lo}\p{M}]+|[^\r\n\p{L}\p{N}]?[\p{Lu}\p{Lt}\p{Lm}\p{Lo}\p{M}]+[\p{Ll}\p{Lm}\p{Lo}\p{M}]*|\p{N}| ?[^\s\p{L}\p{N}]+[\r\n/]*|\s*[\r\n]+|\s+(?!\S)|\s+")
```

[Exact string and validation provenance](../testdata/patterns/P04/manifest.json).

Tokenizer examples:

- [nvidia/NVIDIA-Nemotron-3-Nano-30B-A3B-BF16](https://huggingface.co/nvidia/NVIDIA-Nemotron-3-Nano-30B-A3B-BF16/blob/bf77c3174f68ad409e1c2aa60daeb46e32d1c606/tokenizer.json) — `/pre_tokenizer/pretokenizers/0/pattern/Regex`.
- [swiss-ai/Apertus-8B-Instruct-2509](https://huggingface.co/swiss-ai/Apertus-8B-Instruct-2509/blob/b946d40447b2b597999b9c86d44bee0b452c919f/tokenizer.json) — `/pre_tokenizer/pretokenizers/0/pattern/Regex`.
- [tiiuae/Falcon-H1-7B-Instruct](https://huggingface.co/tiiuae/Falcon-H1-7B-Instruct/blob/41e72f27effbab80cd45b6e884688452253a3686/tokenizer.json) — `/pre_tokenizer/pretokenizers/1/pattern/Regex`.

## P05

```python
regexp(r"(?i:'s|'t|'re|'ve|'m|'ll|'d)|[^\r\n\p{L}\p{N}]?\p{L}+|\p{N}| ?[^\s\p{L}\p{N}]+[\r\n]*|\s*[\r\n]+|\s+(?!\S)|\s+")
```

[Exact string and validation provenance](../testdata/qwen/manifest.json).

Tokenizer examples:

- [Qwen/Qwen2.5-0.5B](https://huggingface.co/Qwen/Qwen2.5-0.5B/blob/060db6499f32faf8b98477b0a26969ef7d8b9987/tokenizer.json) — `/pre_tokenizer/pretokenizers/0/pattern/Regex`.
- [Qwen/Qwen3-8B](https://huggingface.co/Qwen/Qwen3-8B/blob/b968826d9c46dd6066d109eabc6255188de91218/tokenizer.json) — `/pre_tokenizer/pretokenizers/0/pattern/Regex`.
- [Qwen/Qwen3-Coder-30B-A3B-Instruct](https://huggingface.co/Qwen/Qwen3-Coder-30B-A3B-Instruct/blob/b2cff646eb4bb1d68355c01b18ae02e7cf42d120/tokenizer.json) — `/pre_tokenizer/pretokenizers/0/pattern/Regex`.

## P06

```python
regexp(r"[^\r\n\p{L}\p{N}]?[\p{Lu}\p{Lt}\p{Lm}\p{Lo}\p{M}]*[\p{Ll}\p{Lm}\p{Lo}\p{M}]+(?i:'s|'t|'re|'ve|'m|'ll|'d)?|[^\r\n\p{L}\p{N}]?[\p{Lu}\p{Lt}\p{Lm}\p{Lo}\p{M}]+[\p{Ll}\p{Lm}\p{Lo}\p{M}]*(?i:'s|'t|'re|'ve|'m|'ll|'d)?|\p{N}{1,3}| ?[^\s\p{L}\p{N}]+[\r\n/]*|\s*[\r\n]+|\s+(?!\S)|\s+")
```

[Exact string and validation provenance](../testdata/patterns/P06/manifest.json).

Tokenizer examples:

- [MiniMaxAI/MiniMax-M2.1](https://huggingface.co/MiniMaxAI/MiniMax-M2.1/blob/cd97f59135f37b2a6bf09356e485d5e4aeb7dc9c/tokenizer.json) — `/pre_tokenizer/pretokenizers/0/pattern/Regex`.
- [openai/gpt-oss-20b](https://huggingface.co/openai/gpt-oss-20b/blob/6cee5e81ee83917806bbde320786a8fb61efebee/tokenizer.json) — `/pre_tokenizer/pretokenizers/0/pattern/Regex`.
- [microsoft/Phi-4-mini-instruct](https://huggingface.co/microsoft/Phi-4-mini-instruct/blob/cfbefacb99257ffa30c83adab238a50856ac3083/tokenizer.json) — `/pre_tokenizer/pretokenizers/0/pattern/Regex`.

## P07

```python
regexp(r"\p{N}{1,3}")
```

[Exact string and validation provenance](../testdata/patterns/P07/manifest.json).

Tokenizer examples:

- [deepseek-ai/DeepSeek-R1](https://huggingface.co/deepseek-ai/DeepSeek-R1/blob/56d4cbbb4d29f4355bab4b9a39ccb717a14ad5ad/tokenizer.json) — `/pre_tokenizer/pretokenizers/0/pattern/Regex`.
- [stepfun-ai/Step-3.5-Flash](https://huggingface.co/stepfun-ai/Step-3.5-Flash/blob/ab446a3de5e171ea341227e24bb1f090e1b771f7/tokenizer.json) — `/pre_tokenizer/pretokenizers/0/pattern/Regex`.
- [tencent/Hy3](https://huggingface.co/tencent/Hy3/blob/a960ebc3da325ba167f069f76c41eb62c9280d22/tokenizer.json) — `/pre_tokenizer/pretokenizers/0/pattern/Regex`.

## P08

```python
regexp(r"\s?[A-Za-zµÀ-ÖØ-öø-ƺƼ-ƿǄ-ʓʕ-ʯͰ-ͳͶͷͻ-ͽͿΆΈ-ΊΌΎ-ΡΣ-ϵϷ-ҁҊ-ԯԱ-ՖႠ-ჅᎠ-Ᏽᏸ-ᏽᲐ-ᲺᲽ-Ჿᴀ-ᴫᵫ-ᵷᵹ-ᶚḀ-ἕἘ-Ἕἠ-ὅὈ-Ὅὐ-ὗὙὛὝὟ-ώᾀ-ᾴᾶ-ᾼιῂ-ῄῆ-ῌῐ-ΐῖ-Ίῠ-Ῥῲ-ῴῶ-ῼℂℇℊ-ℓℕℙ-ℝℤΩℨK-ℭℯ-ℴℹℼ-ℿⅅ-ⅉⅎↃↄⰀ-ⱻⱾ-ⳤⳫ-ⳮⳲⳳꙀ-ꙭꚀ-ꚛꜢ-ꝯꝱ-ꞇꞋ-ꞎꭰ-ꮿﬀ-ﬆﬓ-ﬗＡ-Ｚａ-ｚ𐐀-𐑏𐒰-𐓓𐓘-𐓻𐲀-𐲲𐳀-𐳲𑢠-𑣟𞤀-𞥃]+")
```

[Exact string and validation provenance](../testdata/patterns/P08/manifest.json).

Tokenizer examples:

- [deepseek-ai/DeepSeek-Coder-V2-Lite-Instruct](https://huggingface.co/deepseek-ai/DeepSeek-Coder-V2-Lite-Instruct/blob/e434a23f91ba5b4923cf6c9d9a238eb4a08e3a11/tokenizer.json) — `/pre_tokenizer/pretokenizers/1/pattern/Regex`.
- [deepseek-ai/DeepSeek-V2](https://huggingface.co/deepseek-ai/DeepSeek-V2/blob/4461458f186c35188585855f28f77af5661ad489/tokenizer.json) — `/pre_tokenizer/pretokenizers/1/pattern/Regex`.
- [deepseek-ai/deepseek-coder-7b-instruct-v1.5](https://huggingface.co/deepseek-ai/deepseek-coder-7b-instruct-v1.5/blob/2a050a4c59d687a85324d32e147517992117ed30/tokenizer.json) — `/pre_tokenizer/pretokenizers/1/pattern/Regex`.

## P09

```python
regexp(r"(?i:'s|'t|'re|'ve|'m|'ll|'d)|[^" + "\r\n" + r"\p{L}\p{N}]?\p{L}+|\p{N}| ?[^\s\p{L}\p{N}]+[" + "\r\n" + r"]*|\s*[" + "\r\n" + r"]+|\s+(?!\S)|\s+")
```

[Exact string and validation provenance](../testdata/patterns/P09/manifest.json).

Tokenizer examples:

- [stabilityai/stablelm-2-1_6b](https://huggingface.co/stabilityai/stablelm-2-1_6b/blob/f499ead74c53749bd93cebc6ce8bc0d7bdf1eaef/tokenizer.json) — `/pre_tokenizer/pretokenizers/0/pattern/Regex`.
- [stabilityai/stablelm-2-1_6b-chat](https://huggingface.co/stabilityai/stablelm-2-1_6b-chat/blob/f3fe67057c2789ae1bb1fe42b038da99840d4f13/tokenizer.json) — `/pre_tokenizer/pretokenizers/0/pattern/Regex`.
- [stabilityai/stablelm-2-zephyr-1_6b](https://huggingface.co/stabilityai/stablelm-2-zephyr-1_6b/blob/2f275b1127d59fc31e4f7c7426d528768ada9ea4/tokenizer.json) — `/pre_tokenizer/pretokenizers/0/pattern/Regex`.

## P10

```python
regexp(r"\A\p{Nd}{1,2}(?=\p{Nd}{3}+\z)")
```

[Exact string and validation provenance](../testdata/patterns/P10/manifest.json).
[Anchor and quantifier semantics](ContextualNumbers.md#p10-semantics).

Tokenizer examples:

- [arcee-ai/Trinity-Large-Thinking](https://huggingface.co/arcee-ai/Trinity-Large-Thinking/blob/dc6a99b6b74202880e31491a445a94884c97890c/tokenizer.json) — `/pre_tokenizer/pretokenizers/2/pattern/Regex`.
- [arcee-ai/Trinity-Mini](https://huggingface.co/arcee-ai/Trinity-Mini/blob/cdb81e8130815fec299cccc8f8700004b8366e1e/tokenizer.json) — `/pre_tokenizer/pretokenizers/2/pattern/Regex`.
- [arcee-ai/Trinity-Nano-Base](https://huggingface.co/arcee-ai/Trinity-Nano-Base/blob/923a95dbba013c78c4fd6dbc6bce31eea553da7c/tokenizer.json) — `/pre_tokenizer/pretokenizers/2/pattern/Regex`.

## P11

```python
regexp(r"[" + "\r\n" + r"]")
```

[Exact string and validation provenance](../testdata/patterns/P11/manifest.json).

Tokenizer examples:

- [deepseek-ai/DeepSeek-Coder-V2-Lite-Instruct](https://huggingface.co/deepseek-ai/DeepSeek-Coder-V2-Lite-Instruct/blob/e434a23f91ba5b4923cf6c9d9a238eb4a08e3a11/tokenizer.json) — `/pre_tokenizer/pretokenizers/0/pattern/Regex`.
- [deepseek-ai/DeepSeek-V2](https://huggingface.co/deepseek-ai/DeepSeek-V2/blob/4461458f186c35188585855f28f77af5661ad489/tokenizer.json) — `/pre_tokenizer/pretokenizers/0/pattern/Regex`.
- [deepseek-ai/deepseek-coder-7b-instruct-v1.5](https://huggingface.co/deepseek-ai/deepseek-coder-7b-instruct-v1.5/blob/2a050a4c59d687a85324d32e147517992117ed30/tokenizer.json) — `/pre_tokenizer/pretokenizers/0/pattern/Regex`.

## P12

```python
regexp(r"(?:[\u4E00-\u9FFF\u3040-\u309F\u30A0-\u30FF\u3400-\u4DBF\uf900-\uFAFF\uFF65-\uFF9F\u2F00-\u2FDF]+|[\u0E00-\u0E7F]+|[\u0E80-\u0EFF]+|[\u1780-\u17FF]+|[\u1000-\u109F\uAA60-\uAA7F\uA9E0-\uA9FF]+|[\uAC00-\uD7AF\u1100-\u11FF]+)")
```

[Exact string and validation provenance](../testdata/patterns/P12/manifest.json).

Tokenizer examples:

- [arcee-ai/Trinity-Large-Thinking](https://huggingface.co/arcee-ai/Trinity-Large-Thinking/blob/dc6a99b6b74202880e31491a445a94884c97890c/tokenizer.json) — `/pre_tokenizer/pretokenizers/4/pattern/Regex`.
- [arcee-ai/Trinity-Mini](https://huggingface.co/arcee-ai/Trinity-Mini/blob/cdb81e8130815fec299cccc8f8700004b8366e1e/tokenizer.json) — `/pre_tokenizer/pretokenizers/4/pattern/Regex`.
- [arcee-ai/Trinity-Nano-Base](https://huggingface.co/arcee-ai/Trinity-Nano-Base/blob/923a95dbba013c78c4fd6dbc6bce31eea553da7c/tokenizer.json) — `/pre_tokenizer/pretokenizers/4/pattern/Regex`.

## P13

```python
regexp(r"'(?i:[sdmt]|ll|ve|re)|[^\r\n\p{L}\p{N}]?+\p{L}+|\p{N}| ?[^\s\p{L}\p{N}]++[\r\n]*|\s*[\r\n]|\s+(?!\S)|\s+")
```

[Exact string and validation provenance](../testdata/patterns/P13/manifest.json).

Tokenizer examples:

- [inclusionAI/LLaDA2.0-mini](https://huggingface.co/inclusionAI/LLaDA2.0-mini/blob/dad945cac317da394b390f82c7b40691d8a881ed/tokenizer.json) — `/pre_tokenizer/pretokenizers/0/pattern/Regex`.
- [inclusionAI/LLaDA2.1-mini](https://huggingface.co/inclusionAI/LLaDA2.1-mini/blob/20e64e2ad21644d0e5248586ed9c942cdd45de0f/tokenizer.json) — `/pre_tokenizer/pretokenizers/0/pattern/Regex`.
- [inclusionAI/Ling-3.0-flash](https://huggingface.co/inclusionAI/Ling-3.0-flash/blob/e0dfe7cd0f6e3b572bbbc0a8a84947469e428cc3/tokenizer.json) — `/pre_tokenizer/pretokenizers/0/pattern/Regex`.

## P15

```python
regexp(r"\p{Hangul}+|[！…“”‘’—：；，、-〿︰-﹏]+|[ㄅ-ㄯ]+|[一-龥぀-ゟ゠-ヿ]+")
```

[Exact string and validation provenance](../testdata/patterns/P15/manifest.json).

Tokenizer examples:

- [tencent/Youtu-LLM-2B](https://huggingface.co/tencent/Youtu-LLM-2B/blob/8b0e73594661a945a7f5aafc737b63e3d8ad3f75/tokenizer.json) — `/pre_tokenizer/pretokenizers/0/pattern/Regex`.

## P16

```python
regexp(r"[0-9]")
```

[Exact string and validation provenance](../testdata/patterns/P16/manifest.json).

Tokenizer examples:

- [tiiuae/Falcon-H1-0.5B-Base](https://huggingface.co/tiiuae/Falcon-H1-0.5B-Base/blob/59fb76e8c5d3fc7441b062be638e1ba0afd5c687/tokenizer.json) — `/pre_tokenizer/pretokenizers/3/pattern/Regex`.
- [tiiuae/Falcon-H1-7B-Instruct](https://huggingface.co/tiiuae/Falcon-H1-7B-Instruct/blob/41e72f27effbab80cd45b6e884688452253a3686/tokenizer.json) — `/pre_tokenizer/pretokenizers/3/pattern/Regex`.

## P17

```python
regexp(r"[一-龥ࠀ-一가-퟿]+")
```

[Exact string and validation provenance](../testdata/patterns/P17/manifest.json).

Tokenizer examples:

- [deepseek-ai/DeepSeek-Coder-V2-Lite-Instruct](https://huggingface.co/deepseek-ai/DeepSeek-Coder-V2-Lite-Instruct/blob/e434a23f91ba5b4923cf6c9d9a238eb4a08e3a11/tokenizer.json) — `/pre_tokenizer/pretokenizers/4/pattern/Regex`.
- [deepseek-ai/DeepSeek-V2](https://huggingface.co/deepseek-ai/DeepSeek-V2/blob/4461458f186c35188585855f28f77af5661ad489/tokenizer.json) — `/pre_tokenizer/pretokenizers/4/pattern/Regex`.
- [deepseek-ai/deepseek-coder-7b-instruct-v1.5](https://huggingface.co/deepseek-ai/deepseek-coder-7b-instruct-v1.5/blob/2a050a4c59d687a85324d32e147517992117ed30/tokenizer.json) — `/pre_tokenizer/pretokenizers/4/pattern/Regex`.

## P18

```python
regexp(r"\s?[!-/:-~！-／：-～‘-‟　-。]+")
```

[Exact string and validation provenance](../testdata/patterns/P18/manifest.json).

Tokenizer examples:

- [deepseek-ai/DeepSeek-Coder-V2-Lite-Instruct](https://huggingface.co/deepseek-ai/DeepSeek-Coder-V2-Lite-Instruct/blob/e434a23f91ba5b4923cf6c9d9a238eb4a08e3a11/tokenizer.json) — `/pre_tokenizer/pretokenizers/2/pattern/Regex`.
- [deepseek-ai/DeepSeek-V2](https://huggingface.co/deepseek-ai/DeepSeek-V2/blob/4461458f186c35188585855f28f77af5661ad489/tokenizer.json) — `/pre_tokenizer/pretokenizers/2/pattern/Regex`.
- [deepseek-ai/deepseek-coder-7b-instruct-v1.5](https://huggingface.co/deepseek-ai/deepseek-coder-7b-instruct-v1.5/blob/2a050a4c59d687a85324d32e147517992117ed30/tokenizer.json) — `/pre_tokenizer/pretokenizers/2/pattern/Regex`.

## P19

```python
regexp(r"[^\r\n\p{L}\p{N}]?[\p{Lu}\p{Lt}\p{Lm}\p{Lo}\p{M}]*[\p{Ll}\p{Lm}\p{Lo}\p{M}]+(?i:'s|'t|'re|'ve|'m|'ll|'d)?|[^\r\n\p{L}\p{N}]?[\p{Lu}\p{Lt}\p{Lm}\p{Lo}\p{M}]+[\p{Ll}\p{Lm}\p{Lo}\p{M}]*(?i:'s|'t|'re|'ve|'m|'ll|'d)?|\p{N}| ?[^\s\p{L}\p{N}]+[\r\n/]*|\s*[\r\n]+|\s+(?!\S)|\s+")
```

[Exact string and validation provenance](../testdata/patterns/P19/manifest.json).

Tokenizer examples:

- [tencent/Youtu-LLM-2B](https://huggingface.co/tencent/Youtu-LLM-2B/blob/8b0e73594661a945a7f5aafc737b63e3d8ad3f75/tokenizer.json) — `/pre_tokenizer/pretokenizers/1/pattern/Regex`.

## P20

```python
regexp(r"(?i:'s|'t|'re|'ve|'m|'ll|'d)|[^\r\n\p{L}\p{N}]?\p{L}+|\p{N}+| ?[^\s\p{L}\p{N}]+[\r\n]*|\s*[\r\n]+|\s+(?!\S)|\s+")
```

[Exact string and validation provenance](../testdata/patterns/P20/manifest.json).

Tokenizer examples:

- [openbmb/MiniCPM5-1B](https://huggingface.co/openbmb/MiniCPM5-1B/blob/87179e5c1f455ef22e6223592d2d61351b525bfc/tokenizer.json) — `/pre_tokenizer/pretokenizers/1/pattern/Regex`.
- [openbmb/MiniCPM5-1B-MLX](https://huggingface.co/openbmb/MiniCPM5-1B-MLX/blob/9879b18bf2928355fcdf4287635388a3665a40cb/tokenizer.json) — `/pre_tokenizer/pretokenizers/1/pattern/Regex`.

## P21

```python
regexp(r"[一-龥぀-ゟ゠-ヿ]+")
```

[Exact string and validation provenance](../testdata/patterns/P21/manifest.json).

Tokenizer examples:

- [deepseek-ai/DeepSeek-R1](https://huggingface.co/deepseek-ai/DeepSeek-R1/blob/56d4cbbb4d29f4355bab4b9a39ccb717a14ad5ad/tokenizer.json) — `/pre_tokenizer/pretokenizers/1/pattern/Regex`.
- [baidu/Unlimited-OCR](https://huggingface.co/baidu/Unlimited-OCR/blob/07dea832e22aefee32ad281d4b80551282e1c168/tokenizer.json) — `/pre_tokenizer/pretokenizers/1/pattern/Regex`.
- [stepfun-ai/Step-3.5-Flash](https://huggingface.co/stepfun-ai/Step-3.5-Flash/blob/ab446a3de5e171ea341227e24bb1f090e1b771f7/tokenizer.json) — `/pre_tokenizer/pretokenizers/1/pattern/Regex`.

## P22

```python
regexp(r"""[!"#$%&'()*+,\-./:;<=>?@\[\\\]^_`{|}~][A-Za-z]+|[^""" + "\r\n" + r"\p{L}\p{P}\p{S}]?[\p{L}\p{M}]+| ?[\p{P}\p{S}]+[" + "\r\n" + r"]*|\s*[" + "\r\n" + r"]+|\s+(?!\S)|\s+")
```

[Exact string and validation provenance](../testdata/patterns/P22/manifest.json).

Tokenizer examples:

- [arcee-ai/Trinity-Large-Thinking](https://huggingface.co/arcee-ai/Trinity-Large-Thinking/blob/dc6a99b6b74202880e31491a445a94884c97890c/tokenizer.json) — `/pre_tokenizer/pretokenizers/5/pattern/Regex`.
- [deepseek-ai/DeepSeek-V3](https://huggingface.co/deepseek-ai/DeepSeek-V3/blob/e815299b0bcbac849fa540c768ef21845365c9eb/tokenizer.json) — `/pre_tokenizer/pretokenizers/2/pattern/Regex`.
- [tencent/Hy3](https://huggingface.co/tencent/Hy3/blob/a960ebc3da325ba167f069f76c41eb62c9280d22/tokenizer.json) — `/pre_tokenizer/pretokenizers/2/pattern/Regex`.

## P23

```python
regexp(r" ?[^(\s|[.,!?…。，、।۔،])]+")
```

[Exact string and validation provenance](../testdata/patterns/P23/manifest.json).

Tokenizer examples:

- [bigscience/bloom](https://huggingface.co/bigscience/bloom/blob/7f10a99ce7c08f03c7719a586cb2cbda1433ac05/tokenizer.json) — `/pre_tokenizer/pretokenizers/0/pattern/Regex`.
- [bigscience/bloom-7b1](https://huggingface.co/bigscience/bloom-7b1/blob/6232703e399354503377bf59dfbb8397fd569e4a/tokenizer.json) — `/pre_tokenizer/pretokenizers/0/pattern/Regex`.
- [bigscience/bloomz-560m](https://huggingface.co/bigscience/bloomz-560m/blob/a2845d7e13dd12efae154a9f1c63fcc2e0cc4b05/tokenizer.json) — `/pre_tokenizer/pretokenizers/0/pattern/Regex`.

## P24

```python
regexp(r"(?i:'s|'t|'re|'ve|'m|'ll|'d)|[^\r\n\p{L}\p{N}]?\p{L}+|\p{N}{1,3}| ?[^\s\p{L}\p{N}]+[\r\n]*|\s*[\r\n]+|\s+(?!\S)|\s+")
```

[Exact string and validation provenance](../testdata/patterns/P24/manifest.json).

Tokenizer examples:

- [HuggingFaceTB/SmolLM3-3B](https://huggingface.co/HuggingFaceTB/SmolLM3-3B/blob/a07cc9a04f16550a088caea529712d1d335b0ac1/tokenizer.json) — `/pre_tokenizer/pretokenizers/0/pattern/Regex`.
- [microsoft/phi-4](https://huggingface.co/microsoft/phi-4/blob/2db69c1c3e91a05d2c64a3185acfbaf36f744e25/tokenizer.json) — `/pre_tokenizer/pretokenizers/0/pattern/Regex`.
- [zai-org/GLM-4.7-Flash](https://huggingface.co/zai-org/GLM-4.7-Flash/blob/7dd20894a642a0aa287e9827cb1a1f7f91386b67/tokenizer.json) — `/pre_tokenizer/pretokenizers/0/pattern/Regex`.
