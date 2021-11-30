package constant

import (
	"fmt"

	"github.com/SpicyChickenFLY/never-todo-cmd/utils/colorful"
)

// Logos
const (
	logoStr = `                   %s____%s
                  %s/   /%s      
        %s____     /   /%s   %s_____     ____  __  __   ____  _____%s
       %s/\   \   /   /%s   %s/' __ \   / ,. \/\ \/\ \ / ,. \/\  __\%s
       %s\ \   \ /   /%s    %s/\ \/\ \ /\  __/\ \ \/ |/\  __/\ \ \_/%s
        %s\ \   v   /%s     %s\ \_\ \_\\ \____\\ \___/\ \____\\ \_\%s
         %s\ \_____/%s       %s\/_/\/_/ \/____/ \/__/  \/____/ \/_/%s
        %s/%s %s\/____/%s            %s__                   __%s
       %s/   ^ \  \%s           %s/\ \___     ____    __\ \     ____%s
      %s/   / \ \__\%s          %s\ \  __\   / __ \  / __\ \   / __ \%s
     %s/   /   \/__/%s           %s\ \ \/__ /\ \_\ \/\ \__\ \ /\ \_\ \%s
    %s/___/%s                     %s\ \____\\ \____/\ \___/\_\\ \____/%s
   %s/___/%s                       %s\/____/ \/___/  \/__/\/_/ \/___/%s`
)

var (
	startMarkGreen = colorful.GetStartMark("default", "default", "green")
	startMarkRed   = colorful.GetStartMark("default", "default", "red")
	endMark        = colorful.GetEndMark()

	// BlackWhiteLogo is not nice
	BlackWhiteLogo = fmt.Sprintf(logoStr,
		"", "",
		"", "",
		"", "", "", "",
		"", "", "", "",
		"", "", "", "",
		"", "", "", "",
		"", "", "", "",
		"", "", "", "", "", "",
		"", "", "", "",
		"", "", "", "",
		"", "", "", "",
		"", "", "", "",
		"", "", "", "",
	)

	// ColorfulLogo is nice
	ColorfulLogo = fmt.Sprintf(logoStr,
		startMarkGreen, endMark,
		startMarkGreen, endMark,
		startMarkGreen, endMark, startMarkRed, endMark,
		startMarkGreen, endMark, startMarkRed, endMark,
		startMarkGreen, endMark, startMarkRed, endMark,
		startMarkGreen, endMark, startMarkRed, endMark,
		startMarkGreen, endMark, startMarkRed, endMark,
		startMarkRed, endMark, startMarkGreen, endMark, startMarkGreen, endMark,
		startMarkRed, endMark, startMarkGreen, endMark,
		startMarkRed, endMark, startMarkGreen, endMark,
		startMarkRed, endMark, startMarkGreen, endMark,
		startMarkRed, endMark, startMarkGreen, endMark,
		startMarkRed, endMark, startMarkGreen, endMark,
	)

	Descirption = `Never todo (CMD): https://github.com/SpicyChickenFLY/never-todo-cmd
use 'never -h' to get help about how to use this`
	Separator = `=======================================================`
)
