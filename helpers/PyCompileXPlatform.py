#!/usr/bin/python3
from prettytable import PrettyTable
import sys
import argparse
import os


def get_args():
    parser = argparse.ArgumentParser()
    parser.add_argument('-o', '--os', type=str, help="Target OS")
    parser.add_argument('-a', '--arch', type=str, help="Compilation Architecture")
    parser.add_argument('-i', '--infile', type=str, help="infile File")
    parser.add_argument('-w', '--write', type=str, help="Write File")
    parser.add_argument('-p', '--print', action='store_true', help="Print OS and Architecture Compatibility")
    return parser.parse_args()


def os_and_archs():
    return {
        "aix": ["ppc64"],
        "android": ["386", "amd64", "arm", "arm64"],
        "darwin": ["386", "amd64", "arm", "arm64"],
        "dragonfly": ["amd64"],
        "freebsd": ["386", "amd64", "arm"],
        "illumos": ["amd64"],
        "js": ["wasm"],
        "linux": ["386", "amd64", "arm", "arm64", "ppc64", "ppc64le", "mips", "mipsle", "mips64", "mips64le", "s390x"],
        "netbsd": ["386", "amd64", "arm"],
        "openbsd": ["386", "amd64", "arm", "arm64"],
        "plan9": ["386", "amd64", "arm"],
        "solaris": ["amd64"],
        "windows": ["386", "amd64"]
    }


def print_archs_table():
    pt = PrettyTable()
    pt.field_names = ["#", "OS", "Supported Architectures"]

    index = 1
    archs = ""
    for os in os_and_archs():
        for arch in os_and_archs()[os]:
            archs += ", " + arch
        pt.add_row([str(index), os, archs.strip()[2:]])
        index += 1
        archs = ""
    pt.align["Supported Architectures"] = "l"
    print(pt)


def print_build():
    pt = PrettyTable()
    pt.field_names = ["Variable", "Detail"]
    pt.add_row(["OS", args.os])
    pt.add_row(["Architecture", args.arch])
    pt.add_row(["Infile", args.infile])
    pt.add_row(["Outfile", args.write])
    pt.align["Detail"] = "l"
    print(pt)


def compile(args):
    
    # Print Build details
    print_build()

    try:
        # Build Go binary without debug info '-s' and dwarf information, '-w'. Reduce file size.
        # Create outfile with infile file minus '.go' extension
        print("[*] Build binary...")
        os.system(f'GOOS={args.os} GOARCH={args.arch} go build -ldflags="-s -w" -o {args.write} {args.infile}')

        # Use UPX to pack binary
        print("[*] UPX pack binary with '--brute'...")
        os.system(f'upx --brute {args.write}')
    except Exception as e:
        print(e)


def exit(msg):
    print(msg)
    sys.exit()


if __name__ == '__main__':
    args = get_args()

    if args.print:
        print_archs_table()

    if args.infile is not None:
        if os.path.exists(args.infile):
            print('[*] File exists')
        else:
            print("[!] Not found. Checking for 'main.go in current directory...")
            if os.path.exists('main.go'):
                print(["[*] main.go found in current directory. We'll build with this..."])
                args.infile = 'main.go'
            else:
                exit('[*] No suitable *.go files found. Check path.')

    if args.os and args.arch and args.infile:
        args.os = args.os.lower().strip()
        args.arch = args.arch.lower().strip()
        args.infile = args.infile.strip()
        if args.write is False:
            args.write = args.infile[0:-3]
            print(f"[!] No outfile specified with '-o' flag. Automatically generated: {args.write}")
        else:
            args.write = args.write.strip()
    else:
        exit("[!] Missing arguments. Run with -h to view help.")

    if args.arch == 'x86' or args.arch == '386' or args.arch == '86':
        args.arch = "386"

    if args.os in os_and_archs().keys() and args.arch in os_and_archs()[args.os]:
        compile(args)
    else:
        print_archs_table()
        exit(f"[!] Invalid OS: {args.os}. Run script with -p flag to view compatibility table.")
