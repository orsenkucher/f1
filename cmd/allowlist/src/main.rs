use core::fmt;
use std::fs;

use regex::Regex;

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let dir = "./groups";
    let parser = GroupParser::new();
    let paths = fs::read_dir(dir)?
        .map(|res| res.map(|e| e.path()))
        .collect::<Result<Vec<_>, _>>()?;
    let mut groups = paths
        .iter()
        .filter(|path| path.is_dir())
        .map(|path| {
            let name = path.file_stem()?;
            let group = parser.parse(name.to_str()?);
            Some(group)
        })
        .collect::<Option<Vec<_>>>()
        .ok_or("can't collect groups")?;
    groups.sort();
    println!("collected {} groups", groups.len());
    let allowlist = groups
        .into_iter()
        .fold(String::from("# Z   A\n"), |mut acc, g| {
            acc.push_str(&g.to_string());
            acc
        });
    fs::write("./allowlist.txt", allowlist)?;
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
        let re = Regex::new(r"^.*?(?P<z>\d+).*?(?P<a>\d+)$").unwrap();
        Self { re }
    }

    fn parse(&self, name: &str) -> Group {
        let caps = self.re.captures(name).unwrap();
        let z: i64 = caps["z"].parse().unwrap();
        let a: i64 = caps["a"].parse().unwrap();
        Group { z, a }
    }
}

// for entry in fs::read_dir(dir)? {
//     let entry = entry?;
//     let path = entry.path();
//     if path.is_dir() {
//         let name = path.file_stem().ok_or("bad path")?;
//         let group = parser.parse(name.to_str().ok_or("bad str")?);
//         println!("Group: {:?}", group)
//     }
// }
