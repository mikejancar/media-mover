# NOTE: Event Source needs to be created for this to work
# [System.Diagnostics.EventLog]::CreateEventSource("MediaMover", "Application")

$output = C:\bin\Get-ProcessOutput.ps1 -FileName "C:\Dev\media-mover\media-mover.exe" -Args """C:\Users\Mikej\Dropbox\Camera Uploads"" ""C:\Media\Pictures"" ""C:\Media\Video\Home Movies"""

if ($output.StandardError) {
  Write-EventLog -LogName "Application" -Source "MediaMover" -EventId 1001 -EntryType "Error" -Message $output.StandardError
}
else {
  Write-EventLog -LogName "Application" -Source "MediaMover" -EventId 1000 -EntryType "Information" -Message $output.StandardOutput
}
