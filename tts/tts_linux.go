package tts

import (
    "io"
    "os/exec"
)


func TtsCommand(message string) *exec.Cmd {
    // need to order them somehow...
    ttsBins := []string{"mimic3","espeak"}
    m := map[string]func(string) (error, *exec.Cmd){
        "mimic3": mimic3,
        "espeak": espeak,
    }
    for _, ttsBin := range ttsBins  {
        _, err := exec.LookPath(ttsBin)
        if err == nil {
            err , cmd := m[ttsBin](message)
            if err != nil {
                continue;
            } else {
                return cmd
            }
        }
    }
    return exec.Command("echo", message)
}


func mimic3(message string) (error, *exec.Cmd) {
    // cmd := fmt.Sprintf("mimic3 %s 2>/dev/null ", message)
    cmd := exec.Command("mimic3", message)
    cmd.Stderr = io.Discard
    return nil, cmd  
}

func espeak(message string) (error, *exec.Cmd) {
    return nil, exec.Command("espeak", message)
}

