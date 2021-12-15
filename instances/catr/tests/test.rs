use catr::Config;

#[test]
fn test_config() {
    let config = Config::new();

    assert_eq!(config.files.len(), 0);
    assert_eq!(config.number_lines, false);
    assert_eq!(config.number_lines(), false);
    assert_eq!(config.number_nonblank_lines, false);
}
