package wx

// #cgo CPPFLAGS: -I${SRCDIR}/windows_386/ -I${SRCDIR}/../wxWidgets/wxWidgets-3.1.0/include -D_FILE_OFFSET_BITS=64 -D__WXMSW__
// #cgo !wxgo_binary_package_build LDFLAGS: -L${SRCDIR}/windows_386/lib -lwxmsw31u -lwxmsw31u_gl -lwxregexu -lwxexpat -lwxtiff -lwxjpeg -lwxpng -lwxzlib -lwxscintilla
// #cgo msys64-workaround LDFLAGS: -Wl,--allow-multiple-definition
// #cgo LDFLAGS: -l:libstdc++.a -Wl,--subsystem,windows -mwindows -lopengl32 -lglu32 -lrpcrt4 -loleaut32 -lole32 -luuid -lwinspool -lwinmm -lshell32 -lshlwapi -lcomctl32 -lcomdlg32 -ladvapi32 -lversion -lwsock32 -lgdi32 -lntdll -lmsvcrt
// #cgo CXXFLAGS: -fpermissive
import "C"
