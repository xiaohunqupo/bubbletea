package tea

// KeyType indicates whether the key is a special key or runes. Special
// keys are things like KeyEnter, KeyBackspace, and so on. Runes keys are just
// regular characters like 'a', '你', 'ض', '🦄', and so on.
//
//	k := Key{Type: KeyRunes, Runes: []rune{'A'}, Mod: ModShift}
//	if k.Type == KeyRunes {
//
//	    fmt.Println(k.Runes)
//	    // Output: A
//
//	    fmt.Println(k.String())
//	    // Output: shift+a
//
//	}
type KeyType int

// Special key symbols.
const (
	// KeyRunes indicates that the key represents rune(s), like 'a', 'b', 'c',
	// and so on.
	KeyRunes KeyType = iota

	// Special names in C0

	KeyBackspace
	KeyTab
	KeyEnter
	KeyEscape

	// Special names in G0

	KeySpace
	KeyDelete

	// Special keys

	KeyUp
	KeyDown
	KeyRight
	KeyLeft
	KeyBegin
	KeyFind
	KeyInsert
	KeySelect
	KeyPgUp
	KeyPgDown
	KeyHome
	KeyEnd

	// Keypad keys

	KeyKpEnter
	KeyKpEqual
	KeyKpMultiply
	KeyKpPlus
	KeyKpComma
	KeyKpMinus
	KeyKpDecimal
	KeyKpDivide
	KeyKp0
	KeyKp1
	KeyKp2
	KeyKp3
	KeyKp4
	KeyKp5
	KeyKp6
	KeyKp7
	KeyKp8
	KeyKp9

	// The following are keys defined in the Kitty keyboard protocol.
	// TODO: Investigate the names of these keys
	KeyKpSep
	KeyKpUp
	KeyKpDown
	KeyKpLeft
	KeyKpRight
	KeyKpPgUp
	KeyKpPgDown
	KeyKpHome
	KeyKpEnd
	KeyKpInsert
	KeyKpDelete
	KeyKpBegin

	// Function keys

	KeyF1
	KeyF2
	KeyF3
	KeyF4
	KeyF5
	KeyF6
	KeyF7
	KeyF8
	KeyF9
	KeyF10
	KeyF11
	KeyF12
	KeyF13
	KeyF14
	KeyF15
	KeyF16
	KeyF17
	KeyF18
	KeyF19
	KeyF20
	KeyF21
	KeyF22
	KeyF23
	KeyF24
	KeyF25
	KeyF26
	KeyF27
	KeyF28
	KeyF29
	KeyF30
	KeyF31
	KeyF32
	KeyF33
	KeyF34
	KeyF35
	KeyF36
	KeyF37
	KeyF38
	KeyF39
	KeyF40
	KeyF41
	KeyF42
	KeyF43
	KeyF44
	KeyF45
	KeyF46
	KeyF47
	KeyF48
	KeyF49
	KeyF50
	KeyF51
	KeyF52
	KeyF53
	KeyF54
	KeyF55
	KeyF56
	KeyF57
	KeyF58
	KeyF59
	KeyF60
	KeyF61
	KeyF62
	KeyF63

	// The following are keys defined in the Kitty keyboard protocol.
	// TODO: Investigate the names of these keys

	KeyCapsLock
	KeyScrollLock
	KeyNumLock
	KeyPrintScreen
	KeyPause
	KeyMenu

	KeyMediaPlay
	KeyMediaPause
	KeyMediaPlayPause
	KeyMediaReverse
	KeyMediaStop
	KeyMediaFastForward
	KeyMediaRewind
	KeyMediaNext
	KeyMediaPrev
	KeyMediaRecord

	KeyLowerVol
	KeyRaiseVol
	KeyMute

	KeyLeftShift
	KeyLeftAlt
	KeyLeftCtrl
	KeyLeftSuper
	KeyLeftHyper
	KeyLeftMeta
	KeyRightShift
	KeyRightAlt
	KeyRightCtrl
	KeyRightSuper
	KeyRightHyper
	KeyRightMeta
	KeyIsoLevel3Shift
	KeyIsoLevel5Shift
)

// Key contains information about a key or release. Keys are always sent to the
// program's update function. There are a couple general patterns you could use
// to check for key presses or releases:
//
//	// Switch on the string representation of the key (shorter)
//	switch msg := msg.(type) {
//	case KeyPressMsg:
//	    switch msg.String() {
//	    case "enter":
//	        fmt.Println("you pressed enter!")
//	    case "a":
//	        fmt.Println("you pressed a!")
//	    }
//	}
//
//	// Switch on the key type (more foolproof)
//	switch msg := msg.(type) {
//	case KeyReleaseMsg:
//	    switch msg.Sym {
//	    case KeyEnter:
//	        fmt.Println("you pressed enter!")
//	    case KeyRunes:
//	        switch string(msg.Runes) {
//	        case "a":
//	            fmt.Println("you pressed a!")
//	        }
//	    }
//	}
//
// Note that Key.Runes will always contain at least one character, so you can
// always safely call Key.Runes[0]. In most cases Key.Runes will only contain
// one character, though certain input method editors (most notably Chinese
// IMEs) can input multiple runes at once.
type Key struct {
	// Runes contains the actual characters received. This usually has a length
	// of 1. Use [Rune()] to get the first key rune received. If the user
	// presses shift+a, the Runes will be `[]rune{'A'}`.
	Runes []rune

	// Type is a special key, like enter, tab, backspace, and so on.
	Type KeyType

	// altRune is the actual, unshifted key pressed by the user. For example,
	// if the user presses shift+a, or caps lock is on, the altRune will be
	// 'a'.
	//
	// In the case of non-latin keyboards, like Arabic, altRune is the
	// unshifted key on the keyboard.
	//
	// This is only available with the Kitty Keyboard Protocol or the Windows
	// Console API.
	altRune rune

	// baseRune is the key pressed according to the standard PC-101 key layout.
	// On internaltional keyboards, this is the key that would be pressed if
	// the keyboard was set to US layout.
	//
	// For example, if the user presses 'q' on a French AZERTY keyboard, the
	// baseRune will be 'q'.
	//
	// This is only available with the Kitty Keyboard Protocol or the Windows
	// Console API.
	baseRune rune

	// Mod is a modifier key, like ctrl, alt, and so on.
	Mod KeyMod

	// IsRepeat indicates whether the key is being held down and sending events
	// repeatedly.
	//
	// This is only available with the Kitty Keyboard Protocol or the Windows
	// Console API.
	IsRepeat bool
}

// KeyPressMsg represents a key press message.
type KeyPressMsg Key

// String implements fmt.Stringer and is quite useful for matching key
// events. For details, on what this returns see [Key.String].
func (k KeyPressMsg) String() string {
	return Key(k).String()
}

// Rune returns the first rune in the Runes field. If the Runes field is empty,
// it returns 0.
func (k KeyPressMsg) Rune() rune {
	return Key(k).Rune()
}

// KeyReleaseMsg represents a key release message.
type KeyReleaseMsg Key

// String implements fmt.Stringer and is quite useful for matching complex key
// events. For details, on what this returns see [Key.String].
func (k KeyReleaseMsg) String() string {
	return Key(k).String()
}

// Rune returns the first rune in the Runes field. If the Runes field is empty,
// it returns 0.
func (k KeyReleaseMsg) Rune() rune {
	return Key(k).Rune()
}

// Rune returns the first rune in the Runes field. If the Runes field is empty,
// it returns 0.
func (k Key) Rune() rune {
	if len(k.Runes) == 0 {
		return 0
	}
	return k.Runes[0]
}

// String implements fmt.Stringer and is used to convert a key to a string.
// While less type safe than looking at the individual fields, it will usually
// be more convenient and readable to use this method when matching against
// keys.
//
// Note that modifier keys are always printed in the following order:
//   - ctrl
//   - alt
//   - shift
//   - meta
//   - hyper
//   - super
//
// For example, you'll always see "ctrl+shift+alt+a" and never
// "shift+ctrl+alt+a".
func (k Key) String() string {
	var s string
	if k.Mod.Contains(ModCtrl) && k.Type != KeyLeftCtrl && k.Type != KeyRightCtrl {
		s += "ctrl+"
	}
	if k.Mod.Contains(ModAlt) && k.Type != KeyLeftAlt && k.Type != KeyRightAlt {
		s += "alt+"
	}
	if k.Mod.Contains(ModShift) && k.Type != KeyLeftShift && k.Type != KeyRightShift {
		s += "shift+"
	}
	if k.Mod.Contains(ModMeta) && k.Type != KeyLeftMeta && k.Type != KeyRightMeta {
		s += "meta+"
	}
	if k.Mod.Contains(ModHyper) && k.Type != KeyLeftHyper && k.Type != KeyRightHyper {
		s += "hyper+"
	}
	if k.Mod.Contains(ModSuper) && k.Type != KeyLeftSuper && k.Type != KeyRightSuper {
		s += "super+"
	}

	runeStr := func(r rune) string {
		// Space is the only invisible printable character.
		if r == ' ' {
			return "space"
		}
		return string(r)
	}
	if k.baseRune != 0 {
		// If a baseRune is present, use it to represent a key using the standard
		// PC-101 key layout.
		s += runeStr(k.baseRune)
	} else if k.altRune != 0 {
		// Otherwise, use the AltRune aka the non-shifted one if present.
		s += runeStr(k.altRune)
	} else if len(k.Runes) > 0 {
		// Else, just print the rune.
		if len(k.Runes) > 1 {
			s += string(k.Runes)
		} else {
			s += runeStr(k.Rune())
		}
	} else {
		s += k.Type.String()
	}
	return s
}

// String returns the string representation of the key type.
func (k KeyType) String() string {
	if s, ok := keyTypeString[k]; ok {
		return s
	}
	return ""
}

var keyTypeString = map[KeyType]string{
	KeyRunes:      "runes",
	KeyEnter:      "enter",
	KeyTab:        "tab",
	KeyBackspace:  "backspace",
	KeyEscape:     "esc",
	KeySpace:      "space",
	KeyUp:         "up",
	KeyDown:       "down",
	KeyLeft:       "left",
	KeyRight:      "right",
	KeyBegin:      "begin",
	KeyFind:       "find",
	KeyInsert:     "insert",
	KeyDelete:     "delete",
	KeySelect:     "select",
	KeyPgUp:       "pgup",
	KeyPgDown:     "pgdown",
	KeyHome:       "home",
	KeyEnd:        "end",
	KeyKpEnter:    "kpenter",
	KeyKpEqual:    "kpequal",
	KeyKpMultiply: "kpmul",
	KeyKpPlus:     "kpplus",
	KeyKpComma:    "kpcomma",
	KeyKpMinus:    "kpminus",
	KeyKpDecimal:  "kpperiod",
	KeyKpDivide:   "kpdiv",
	KeyKp0:        "kp0",
	KeyKp1:        "kp1",
	KeyKp2:        "kp2",
	KeyKp3:        "kp3",
	KeyKp4:        "kp4",
	KeyKp5:        "kp5",
	KeyKp6:        "kp6",
	KeyKp7:        "kp7",
	KeyKp8:        "kp8",
	KeyKp9:        "kp9",

	// Kitty keyboard extension
	KeyKpSep:    "kpsep",
	KeyKpUp:     "kpup",
	KeyKpDown:   "kpdown",
	KeyKpLeft:   "kpleft",
	KeyKpRight:  "kpright",
	KeyKpPgUp:   "kppgup",
	KeyKpPgDown: "kppgdown",
	KeyKpHome:   "kphome",
	KeyKpEnd:    "kpend",
	KeyKpInsert: "kpinsert",
	KeyKpDelete: "kpdelete",
	KeyKpBegin:  "kpbegin",

	KeyF1:  "f1",
	KeyF2:  "f2",
	KeyF3:  "f3",
	KeyF4:  "f4",
	KeyF5:  "f5",
	KeyF6:  "f6",
	KeyF7:  "f7",
	KeyF8:  "f8",
	KeyF9:  "f9",
	KeyF10: "f10",
	KeyF11: "f11",
	KeyF12: "f12",
	KeyF13: "f13",
	KeyF14: "f14",
	KeyF15: "f15",
	KeyF16: "f16",
	KeyF17: "f17",
	KeyF18: "f18",
	KeyF19: "f19",
	KeyF20: "f20",
	KeyF21: "f21",
	KeyF22: "f22",
	KeyF23: "f23",
	KeyF24: "f24",
	KeyF25: "f25",
	KeyF26: "f26",
	KeyF27: "f27",
	KeyF28: "f28",
	KeyF29: "f29",
	KeyF30: "f30",
	KeyF31: "f31",
	KeyF32: "f32",
	KeyF33: "f33",
	KeyF34: "f34",
	KeyF35: "f35",
	KeyF36: "f36",
	KeyF37: "f37",
	KeyF38: "f38",
	KeyF39: "f39",
	KeyF40: "f40",
	KeyF41: "f41",
	KeyF42: "f42",
	KeyF43: "f43",
	KeyF44: "f44",
	KeyF45: "f45",
	KeyF46: "f46",
	KeyF47: "f47",
	KeyF48: "f48",
	KeyF49: "f49",
	KeyF50: "f50",
	KeyF51: "f51",
	KeyF52: "f52",
	KeyF53: "f53",
	KeyF54: "f54",
	KeyF55: "f55",
	KeyF56: "f56",
	KeyF57: "f57",
	KeyF58: "f58",
	KeyF59: "f59",
	KeyF60: "f60",
	KeyF61: "f61",
	KeyF62: "f62",
	KeyF63: "f63",

	// Kitty keyboard extension
	KeyCapsLock:         "capslock",
	KeyScrollLock:       "scrolllock",
	KeyNumLock:          "numlock",
	KeyPrintScreen:      "printscreen",
	KeyPause:            "pause",
	KeyMenu:             "menu",
	KeyMediaPlay:        "mediaplay",
	KeyMediaPause:       "mediapause",
	KeyMediaPlayPause:   "mediaplaypause",
	KeyMediaReverse:     "mediareverse",
	KeyMediaStop:        "mediastop",
	KeyMediaFastForward: "mediafastforward",
	KeyMediaRewind:      "mediarewind",
	KeyMediaNext:        "medianext",
	KeyMediaPrev:        "mediaprev",
	KeyMediaRecord:      "mediarecord",
	KeyLowerVol:         "lowervol",
	KeyRaiseVol:         "raisevol",
	KeyMute:             "mute",
	KeyLeftShift:        "leftshift",
	KeyLeftAlt:          "leftalt",
	KeyLeftCtrl:         "leftctrl",
	KeyLeftSuper:        "leftsuper",
	KeyLeftHyper:        "lefthyper",
	KeyLeftMeta:         "leftmeta",
	KeyRightShift:       "rightshift",
	KeyRightAlt:         "rightalt",
	KeyRightCtrl:        "rightctrl",
	KeyRightSuper:       "rightsuper",
	KeyRightHyper:       "righthyper",
	KeyRightMeta:        "rightmeta",
	KeyIsoLevel3Shift:   "isolevel3shift",
	KeyIsoLevel5Shift:   "isolevel5shift",
}
