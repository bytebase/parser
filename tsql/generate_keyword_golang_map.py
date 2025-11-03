from typing import List


def extract_keyword_candidates(lines: List[str]) -> List[str]:
    result = []
    for line in lines:
        line = line.strip()
        if not line.startswith(":") and not line.startswith("|"):
            continue
        line = line[1:].strip()
        buf = ""
        for i in range(len(line)):
            if line[i].isdigit() or line[i].isalpha() or line[i] == "_":
                buf += line[i]
                continue
            break
        if len(buf) > 0:
            result.append(buf)
            buf = ""
    return result


        
if __name__ == "__main__":
    with open("TSqlParser.g4") as f:
        content = f.read()
        content_lines = content.split("\n")
        for (idx, line) in enumerate(content_lines):
            if not line.startswith("keyword"):
                continue

            if not (idx < len(content_lines) - 1 and content_lines[idx + 1].endswith(": ABORT")):
                continue
            
            buffer = []
            for j in range(idx + 1, len(content_lines)):
                if content_lines[j].endswith(";"):
                    break
                buffer.append(content_lines[j])
        with open("keyword_map.go", "w") as f:
            f.write(r"// Code generated from generate_keyword_golang_map.py. DO NOT EDIT")
            f.write("\n\n")
            f.write(r"package parser")
            f.write("\n\n")
            f.write(r"var keywordMap map[int]struct{} = map[int]struct{}{")
            candidates = extract_keyword_candidates(buffer)
            for (_, candidate) in enumerate(candidates):
                f.write("\n")
                f.write(" "*4)
                f.write(f"TSqlLexer{candidate}:")
                f.write(" {},") 
            f.write("\n")
            f.write("}")
        print("Write keyword map to keyword_map.go successfully!")

