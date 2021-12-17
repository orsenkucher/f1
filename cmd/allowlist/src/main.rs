use core::fmt;
use std::{fs, path::PathBuf};

use regex::Regex;

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let paths = paths("./groups")?;
    let mut groups = groups(paths).ok_or("can't collect groups")?;
    groups.sort();
    println!("collected {} groups", groups.len());
    write_list(groups)?;
    Ok(())
}

fn paths(dir: &str) -> std::io::Result<Vec<PathBuf>> {
    fs::read_dir(dir)?
        .map(|res| res.map(|e| e.path()))
        .collect::<Result<Vec<_>, _>>()
}

fn groups(paths: Vec<PathBuf>) -> Option<Vec<Group>> {
    let parser = GroupParser::new();
    paths
        .into_iter()
        .filter(|p| p.is_dir())
        .map(|p| Some(parser.parse(p.file_stem()?.to_str()?)))
        .collect::<Option<Vec<_>>>()
}

fn write_list(groups: Vec<Group>) -> std::io::Result<()> {
    let text = groups
        .into_iter()
        .fold(String::from("# Z   A\n"), |mut acc, g| {
            acc.push_str(&g.to_string());
            acc
        });
    fs::write("./allowlist.txt", text)?;
    Ok(())
}

#[derive(Debug, PartialEq, Eq, PartialOrd, Ord)]
struct Group {
    z: i64,
    a: i64,
}

impl fmt::Display for Group {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        let Self { z, a } = self;
        writeln!(f, "{z:4} {a:3}")
    }
}

struct GroupParser {
    re: Regex,
}

impl GroupParser {
    fn new() -> Self {
        let re = Regex::new(r"^.*?(?P<z>\d+).*?(?P<a>\d+)$").expect("regex to comile");
        Self { re }
    }

    fn parse(&self, name: &str) -> Group {
        let caps = self
            .re
            .captures(name)
            .expect("group to have atomic number and mass");
        let z: i64 = caps["z"].parse().expect("to parse Z value as i64");
        let a: i64 = caps["a"].parse().expect("to parse A value as i64");
        Group { z, a }
    }
}
