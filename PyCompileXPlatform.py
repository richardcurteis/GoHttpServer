#!/usr/bin/python3
from prettytable import PrettyTable
import sys
import argparse
import os

def get_args():
	parser = argparse.ArgumentParser()
	parser.add_argument('-o', '--os', type=str, help="Target OS")
	parser.add_argument('-a', '--arch', type=str, help="Compilation Architecture")
	parser.add_argument('-i', '--input', type=str, help="Input File")
	parser.add_argument('-p', '--print', action='store_true', help="Print OS and Architecture Compatibility")
	return parser.parse_args()

def os_and_archs():
	return {
		"aix": ["ppc64"],
		"android": ["386/x86", "amd64", "arm", "arm64"],
		"darwin":	["386/x86", "amd64", "arm", "arm64"],
		"dragonfly": ["amd64"],
		"freebsd": ["386/x86", "amd64", "arm"],
		"illumos": ["amd64"],
		"js": ["wasm"],
		"linux": ["386/x86", "amd64", "arm", "arm64", "ppc64", "ppc64le", "mips", "mipsle", "mips64", "mips64le", "s390x"],
		"netbsd": ["386/x86", "amd64", "arm"],
		"openbsd": ["386/x86", "amd64", "arm", "arm64"],
		"plan9": ["386/x86", "amd64", "arm"],
		"solaris": ["amd64"],
		"windows": ["386/x86", "amd64"]
	}

def print_table():
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

def compile(args):
	os = args.os
	if os is "386/x86":
		os = '386'

	os.system('export' + f'GOARCH={args.arch}')
	os.system('export' + f'GOOS={os}')

	try:
		os.system('go' + 'build' + '-ldflags="-s -w"')
		os.system('upx' + 'brute' + f'{args.input}')
	except Exception as e:
		print(e)
		
def exit(msg):
	print(msg)
	sys.exit()

if __name__ == '__main__':
	args = get_args()

	if args.print:
		print_table()

	if args.os and args.arch and args.input:
		args.os = args.os.lower().strip()
		args.arch = args.arch.lower().strip()
		args.input = args.input.strip()
	else:
		exit("[!] Missing arguments. Run with -h to view help.")
	
	if args.arch is 'x86' or '386' or '86':
		args.arch = "386/x86"
	
	if os.path.exists(args.input):
		print('[*] File exists')
	else:
		exit('[*] File not found. Check path.')

	if args.os in os_and_archs().keys() and args.arch in os_and_archs()[args.os]:
		compile(args)
	else:
		print_table()
		exit(f"[!] Invalid OS: {args.os}. Run script with -p flag to view compatibility table.")
