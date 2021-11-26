// Code generated by "stringer -type=Type"; DO NOT EDIT.

package fileid

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Thumbnail-0]
	_ = x[ProfilePhoto-1]
	_ = x[Photo-2]
	_ = x[Voice-3]
	_ = x[Video-4]
	_ = x[Document-5]
	_ = x[Encrypted-6]
	_ = x[Temp-7]
	_ = x[Sticker-8]
	_ = x[Audio-9]
	_ = x[Animation-10]
	_ = x[EncryptedThumbnail-11]
	_ = x[Wallpaper-12]
	_ = x[VideoNote-13]
	_ = x[SecureRaw-14]
	_ = x[Secure-15]
	_ = x[Background-16]
	_ = x[DocumentAsFile-17]
	_ = x[lastType-18]
}

const _Type_name = "ThumbnailProfilePhotoPhotoVoiceVideoDocumentEncryptedTempStickerAudioAnimationEncryptedThumbnailWallpaperVideoNoteSecureRawSecureBackgroundDocumentAsFilelastType"

var _Type_index = [...]uint8{0, 9, 21, 26, 31, 36, 44, 53, 57, 64, 69, 78, 96, 105, 114, 123, 129, 139, 153, 161}

func (i Type) String() string {
	if i < 0 || i >= Type(len(_Type_index)-1) {
		return "Type(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}