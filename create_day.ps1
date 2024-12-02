param (
    [Parameter(Mandatory = $true)]
    [int]$DayNumber
)

$year = 2024
$folderName = "day$DayNumber"

$goFilePath = "$folderName\day$DayNumber.go"
$testFilePath = "$folderName\day${DayNumber}_test.go"
$inputFilePath = "input\$DayNumber.txt"
$mainFilePath = "main.go"

Get-Content .env | ForEach-Object {
	$name, $value = $_.split('=')
	New-Variable -Name $name -Value $value
}

New-Item -ItemType Directory -Path $folderName -Force | Out-Null

$goFileContent = @"
package $folderName

import (
	"fmt"

	"github.com/jeroen-plug/advent-of-code-$year/input"
)

func Day$DayNumber() {
	lines := input.Lines($DayNumber)

	fmt.Printf("day ${DayNumber}a: %d\n", day${DayNumber}a(lines))
	fmt.Printf("day ${DayNumber}b: %d\n", day${DayNumber}b(lines))
}

func day${DayNumber}a(lines []string) int {
    // TODO: Implement day${DayNumber}a
    return 0
}

func day${DayNumber}b(lines []string) int {
    // TODO: Implement day${DayNumber}b
    return 0
}
"@

if (-Not (Test-Path -Path $goFilePath)) {
    Set-Content -Path $goFilePath -Value $goFileContent
}

$testFileContent = @"
package $folderName

import (
	"strings"
	"testing"
)

const example = ````

func TestDay${DayNumber}a(t *testing.T) {
	want := 0
	res := day${DayNumber}a(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day${DayNumber}a() = %d, want %d", res, want)
	}
}

func TestDay${DayNumber}b(t *testing.T) {
	want := 0
	res := day${DayNumber}b(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day${DayNumber}b() = %d, want %d", res, want)
	}
}
"@

if (-Not (Test-Path -Path $testFilePath)) {
    Set-Content -Path $testFilePath -Value $testFileContent
}

if (-Not (Test-Path -Path $inputFilePath)) {
	$session = New-Object Microsoft.PowerShell.Commands.WebRequestSession
	$cookie = New-Object System.Net.Cookie('session', $SESSION_COOKIE, '/', 'adventofcode.com')
	$session.Cookies.Add($cookie)

	try {
		Invoke-WebRequest -Uri "https://adventofcode.com/$year/day/$DayNumber/input" -WebSession $session -OutFile $inputFilePath
	} catch {
		Write-Host -ForegroundColor Yellow "Warning: Could not download input for day $DayNumber!"
		Set-Content -Path $inputFilePath -Value ""
	}
}

$newImportLine = "	""github.com/jeroen-plug/advent-of-code-$year/day$DayNumber"""
$newCaseLine = @"
	case ${DayNumber}:
        day$DayNumber.Day$DayNumber()
"@

$mainFileContent = Get-Content -Path $mainFilePath -Raw
if ($mainFileContent -notmatch "case ${DayNumber}:") {
	$updatedContent = $mainFileContent -replace "(\s+default:)", "`n$newCaseLine$&"
	Set-Content -Path $mainFilePath -Value $updatedContent
}

$mainFileContent = Get-Content -Path $mainFilePath -Raw
if ($mainFileContent -notmatch [regex]::Escape($newImportLine)) {
	$updatedContent = $mainFileContent -replace "(\n\))", "`n$newImportLine$&"
	Set-Content -Path $mainFilePath -Value $updatedContent
}
