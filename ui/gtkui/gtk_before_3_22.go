//+build gtk_3_6 gtk_3_8 gtk_3_10 gtk_3_12 gtk_3_14 gtk_3_16 gtk_3_18 gtk_3_20

package gtkui

import "github.com/d2r2/gotk3/gtk"

// SetScrolledWindowPropogatedHeight compiled for GTK+ before 3.22 does nothing.
func SetScrolledWindowPropogatedHeight(sw *gtk.ScrolledWindow, propagate bool) {
	// No call
}
