module decodeur

go 1.24

toolchain go1.24.1

replace local/structures => ../StructuresCras

require local/structures v0.0.0

replace local/goFonctions => ../goFonctions

require local/goFonctions v0.0.0

require github.com/google/certificate-transparency-go v1.3.1