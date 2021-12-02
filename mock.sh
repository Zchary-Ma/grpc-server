function gen_go_mock() {
    echo   "Generating go mock"
    mockgen \
      --package=mock . NoteServiceClient \
      > ../mock/note_mock.go
}

gen_go_mock