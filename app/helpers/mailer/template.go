package mailer

import (
	"fmt"
	"os"
)

var (
	//LinkAssets ...
	LinkAssets = os.Getenv("EMAIL_ASSET_CDN")
	//EmailLink ...
	EmailLink = os.Getenv("EMAIL_LINK")
	//EmailLabel ...
	EmailLabel = os.Getenv("EMAIL_LABELS")
)

//Template ...
type Template interface {
	Compose() string
	ComposeForgotPass() string
}
type template struct {
	title       string
	logo        string
	image       string
	foreword    string
	content     string
	buttonLink  string
	buttonLabel string
	footer      string
	logoI       string
	mailLink    string
	mailLabel   string
	message     string
}

func (t *template) setAsset() {
	link := func(asset string) string {
		return fmt.Sprintf("%v/%v", LinkAssets, asset)
	}
	t.logo = link("logo.png")
	t.image = link("image.png")
	t.logoI = link("logo_white.png")
	t.mailLink = EmailLink
	t.mailLabel = EmailLabel
}

func (t *template) ComposeForgotPass() string {
	t.setAsset()
	return fmt.Sprintf(`<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional //EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd"><html xmlns="http://www.w3.org/1999/xhtml" xmlns:o="urn:schemas-microsoft-com:office:office" xmlns:v="urn:schemas-microsoft-com:vml"><head><meta content="text/html; charset=utf-8" http-equiv="Content-Type"/><meta content="width=device-width" name="viewport"/><meta content="IE=edge" http-equiv="X-UA-Compatible"/><title>%v</title><style type="text/css">body{margin: 0;padding: 0;}table,td,tr{vertical-align: top;border-collapse: collapse;}*{line-height: inherit;}a[x-apple-data-detectors=true]{color: inherit !important;text-decoration: none !important;}</style><style id="media-query" type="text/css">@media (max-width: 720px){.block-grid,.col{min-width: 320px !important;max-width: 100%% !important;display: block !important;}.block-grid{width: 100%% !important;}.col{width: 100%% !important;}.col_cont{margin: 0 auto;}img.fullwidth,img.fullwidthOnMobile{max-width: 100%% !important;}.no-stack .col{min-width: 0 !important;display: table-cell !important;}.no-stack.two-up .col{width: 50%% !important;}.no-stack .col.num2{width: 16.6%% !important;}.no-stack .col.num3{width: 25%% !important;}.no-stack .col.num4{width: 33%% !important;}.no-stack .col.num5{width: 41.6%% !important;}.no-stack .col.num6{width: 50%% !important;}.no-stack .col.num7{width: 58.3%% !important;}.no-stack .col.num8{width: 66.6%% !important;}.no-stack .col.num9{width: 75%% !important;}.no-stack .col.num10{width: 83.3%% !important;}.video-block{max-width: none !important;}.mobile_hide{min-height: 0px;max-height: 0px;max-width: 0px;display: none;overflow: hidden;font-size: 0px;}.desktop_hide{display: block !important;max-height: none !important;}}</style></head><body class="clean-body" style="margin: 0; padding: 0; -webkit-text-size-adjust: 100%%; background-color: #FFFFFF;"><table bgcolor="#FFFFFF" cellpadding="0" cellspacing="0" class="nl-container" role="presentation" style="table-layout: fixed; vertical-align: top; min-width: 320px; border-spacing: 0; border-collapse: collapse; mso-table-lspace: 0pt; mso-table-rspace: 0pt; background-color: #FFFFFF; width: 100%%;" valign="top" width="100%%"><tbody><tr style="vertical-align: top;" valign="top"><td style="word-break: break-word; vertical-align: top;" valign="top"><div style="background-color:transparent;"><div class="block-grid mixed-two-up" style="min-width: 320px; max-width: 700px; overflow-wrap: break-word; word-wrap: break-word; word-break: break-word; Margin: 0 auto; background-color: transparent;"><div style="border-collapse: collapse;display: table;width: 100%%;background-color:transparent;"><div class="col num3" style="display: table-cell; vertical-align: top; max-width: 320px; min-width: 174px; width: 175px;"><div class="col_cont" style="width:100%% !important;"><div style="border-top:0px solid transparent; border-left:0px solid transparent; border-bottom:0px solid transparent; border-right:0px solid transparent; padding-top:5px; padding-bottom:5px; padding-right: 0px; padding-left: 0px;"><div align="left" class="img-container left autowidth" style="padding-right: 0px;padding-left: 10px;"><div style="font-size:1px;line-height:5px"></div><img alt="Digisoul Logo" border="0" class="left autowidth" src="%v" style="text-decoration: none; -ms-interpolation-mode: bicubic; height: auto; border: 0; width: 100%%; max-width: 117px; display: block;" title="Alternate text" width="117"/><div style="font-size:1px;line-height:5px"></div></div></div></div></div><div class="col num9" style="display: table-cell; vertical-align: top; max-width: 320px; min-width: 522px; width: 525px;"><div class="col_cont" style="width:100%% !important;"><div style="border-top:0px solid transparent; border-left:0px solid transparent; border-bottom:0px solid transparent; border-right:0px solid transparent; padding-top:5px; padding-bottom:5px; padding-right: 0px; padding-left: 0px;"><div></div></div></div></div></div></div></div><div style="background-color:transparent;"><div class="block-grid" style="min-width: 320px; max-width: 700px; overflow-wrap: break-word; word-wrap: break-word; word-break: break-word; Margin: 0 auto; background-color: transparent;"><div style="border-collapse: collapse;display: table;width: 100%%;background-color:transparent;"><div class="col num12" style="min-width: 320px; max-width: 700px; display: table-cell; vertical-align: top; width: 700px;"><div class="col_cont" style="width:100%% !important;"><div style="border-top:0px solid transparent; border-left:0px solid transparent; border-bottom:0px solid transparent; border-right:0px solid transparent; padding-top:5px; padding-bottom:5px; padding-right: 0px; padding-left: 0px;"><div style="color:#555555;font-family:Roboto, Tahoma, Verdana, Segoe, sans-serif;line-height:1.2;padding-top:15px;padding-right:10px;padding-bottom:15px;padding-left:10px;"><div style="line-height: 1.2; font-size: 12px; color: #555555; font-family: Roboto, Tahoma, Verdana, Segoe, sans-serif; mso-line-height-alt: 14px;"><p style="font-size: 14px; line-height: 1.2; word-break: break-word; mso-line-height-alt: 17px; margin: 0;"><strong>%v</strong></p></div></div></div></div></div></div></div></div><div style="background-color:transparent;"><div class="block-grid" style="min-width: 320px; max-width: 700px; overflow-wrap: break-word; word-wrap: break-word; word-break: break-word; Margin: 0 auto; background-color: transparent;"><div style="border-collapse: collapse;display: table;width: 100%%;background-color:transparent;"><div class="col num12" style="min-width: 320px; max-width: 700px; display: table-cell; vertical-align: top; width: 700px;"><div class="col_cont" style="width:100%% !important;"><div style="border-top:0px solid transparent; border-left:0px solid transparent; border-bottom:0px solid transparent; border-right:0px solid transparent; padding-top:5px; padding-bottom:5px; padding-right: 0px; padding-left: 0px;"><div style="color:#555555;font-family:Roboto, Tahoma, Verdana, Segoe, sans-serif;line-height:1.2;padding-top:10px;padding-right:10px;padding-bottom:10px;padding-left:10px;"><div style="line-height: 1.2; font-size: 12px; color: #555555; font-family: Roboto, Tahoma, Verdana, Segoe, sans-serif; mso-line-height-alt: 14px;"><p style="font-size: 14px; line-height: 1.2; word-break: break-word; mso-line-height-alt: 17px; margin: 0;">%v</p></div></div></div></div></div></div></div></div><div style="background-color:transparent;"><div class="block-grid" style="min-width: 320px; max-width: 700px; overflow-wrap: break-word; word-wrap: break-word; word-break: break-word; Margin: 0 auto; background-color: transparent;"><div style="border-collapse: collapse;display: table;width: 100%%;background-color:transparent;"><div class="col num12" style="min-width: 320px; max-width: 700px; display: table-cell; vertical-align: top; width: 700px;"><div class="col_cont" style="width:100%% !important;"><div style="border-top:0px solid transparent; border-left:0px solid transparent; border-bottom:0px solid transparent; border-right:0px solid transparent; padding-top:5px; padding-bottom:5px; padding-right: 0px; padding-left: 0px;"><div style="color:#555555;font-family:Roboto, Tahoma, Verdana, Segoe, sans-serif;line-height:1.2;padding-top:10px;padding-right:10px;padding-bottom:10px;padding-left:10px;"><div style="line-height: 1.2; font-size: 12px; color: #555555; font-family: Roboto, Tahoma, Verdana, Segoe, sans-serif; mso-line-height-alt: 14px;"><p style="font-size: 14px; line-height: 1.2; word-break: break-word; mso-line-height-alt: 17px; margin: 0;">Password Baru: %v</p></div></div></div></div></div></div></div></div><div style="background-color:transparent;"><div class="block-grid" style="min-width: 320px; max-width: 700px; overflow-wrap: break-word; word-wrap: break-word; word-break: break-word; Margin: 0 auto; background-color: transparent;"><div style="border-collapse: collapse;display: table;width: 100%%;background-color:transparent;"><div class="col num12" style="min-width: 320px; max-width: 700px; display: table-cell; vertical-align: top; width: 700px;"><div class="col_cont" style="width:100%% !important;"><div style="border-top:0px solid transparent; border-left:0px solid transparent; border-bottom:0px solid transparent; border-right:0px solid transparent; padding-top:5px; padding-bottom:5px; padding-right: 0px; padding-left: 0px;"><div style="color:#555555;font-family:Roboto, Tahoma, Verdana, Segoe, sans-serif;line-height:1.2;padding-top:10px;padding-right:10px;padding-bottom:15px;padding-left:10px;"><div style="line-height: 1.2; font-size: 12px; color: #555555; font-family: Roboto, Tahoma, Verdana, Segoe, sans-serif; mso-line-height-alt: 14px;"><p style="font-size: 14px; line-height: 1.2; word-break: break-word; mso-line-height-alt: 17px; margin: 0;">%v</p></div></div></div></div></div></div></div></div><div style="background-color:transparent;"><div class="block-grid mixed-two-up" style="min-width: 320px; max-width: 700px; overflow-wrap: break-word; word-wrap: break-word; word-break: break-word; Margin: 0 auto; background-color: transparent;"><div style="border-collapse: collapse;display: table;width: 100%%;background-color:transparent;"><div class="col num4" style="display: table-cell; vertical-align: top; max-width: 320px; min-width: 232px; width: 233px;"><div class="col_cont" style="width:100%% !important;"><div style="border-top:0px solid transparent; border-left:0px solid transparent; border-bottom:0px solid transparent; border-right:0px solid transparent; padding-top:5px; padding-bottom:5px; padding-right: 0px; padding-left: 0px;"><div style="color:#555555;font-family:Roboto, Tahoma, Verdana, Segoe, sans-serif;line-height:1.2;padding-top:10px;padding-right:10px;padding-bottom:25px;padding-left:10px;"><div style="line-height: 1.2; font-size: 12px; color: #555555; font-family: Roboto, Tahoma, Verdana, Segoe, sans-serif; mso-line-height-alt: 14px;"><p style="font-size: 14px; line-height: 1.2; word-break: break-word; mso-line-height-alt: 17px; margin: 0;">Salam,</p></div></div><div style="color:#555555;font-family:Roboto, Tahoma, Verdana, Segoe, sans-serif;line-height:1.2;padding-top:30px;padding-right:10px;padding-bottom:15px;padding-left:10px;"><div style="line-height: 1.2; font-size: 12px; color: #555555; font-family: Roboto, Tahoma, Verdana, Segoe, sans-serif; mso-line-height-alt: 14px;"><p style="line-height: 1.2; word-break: break-word; font-size: 14px; mso-line-height-alt: 17px; margin: 0;"><span style="font-size: 14px;">Tim StoryTale</span></p></div></div></div></div></div><div class="col num8" style="display: table-cell; vertical-align: top; max-width: 320px; min-width: 464px; width: 466px;"><div class="col_cont" style="width:100%% !important;"><div style="border-top:0px solid transparent; border-left:0px solid transparent; border-bottom:0px solid transparent; border-right:0px solid transparent; padding-top:5px; padding-bottom:5px; padding-right: 0px; padding-left: 0px;"><div></div></div></div></div></div></div></div><div style="background-color:transparent;"><div class="block-grid mixed-two-up" style="min-width: 320px; max-width: 700px; overflow-wrap: break-word; word-wrap: break-word; word-break: break-word; Margin: 0 auto; background-color: #575fce;"><div style="border-collapse: collapse;display: table;width: 100%%;background-color:#575fce;"><div class="col num4" style="display: table-cell; vertical-align: top; max-width: 320px; min-width: 232px; width: 233px;"><div class="col_cont" style="width:100%% !important;"><div style="border-top:0px solid transparent; border-left:0px solid transparent; border-bottom:0px solid transparent; border-right:0px solid transparent; padding-top:5px; padding-bottom:5px; padding-right: 0px; padding-left: 0px;"><div align="left" class="img-container left autowidth" style="padding-right: 0px;padding-left: 15px;"><div style="font-size:1px;line-height:5px"></div><img alt="Alternate text" border="0" class="left autowidth" src="%v" style="text-decoration: none; -ms-interpolation-mode: bicubic; height: auto; border: 0; width: 100%%; max-width: 116px; display: block;" title="Alternate text" width="116"/></div></div></div></div><div class="col num8" style="display: table-cell; vertical-align: top; max-width: 320px; min-width: 464px; width: 466px;"><div class="col_cont" style="width:100%% !important;"><div style="border-top:0px solid transparent; border-left:0px solid transparent; border-bottom:0px solid transparent; border-right:0px solid transparent; padding-top:5px; padding-bottom:5px; padding-right: 0px; padding-left: 0px;"><div></div></div></div></div></div></div></div><div style="background-color:transparent;"><div class="block-grid" style="min-width: 320px; max-width: 700px; overflow-wrap: break-word; word-wrap: break-word; word-break: break-word; Margin: 0 auto; background-color: #575fce;"><div style="border-collapse: collapse;display: table;width: 100%%;background-color:#575fce;"><div class="col num12" style="min-width: 320px; max-width: 700px; display: table-cell; vertical-align: top; width: 700px;"><div class="col_cont" style="width:100%% !important;"><div style="border-top:0px solid transparent; border-left:0px solid transparent; border-bottom:0px solid transparent; border-right:0px solid transparent; padding-top:5px; padding-bottom:5px; padding-right: 0px; padding-left: 0px;"><div style="color:#ffffff;font-family:Roboto, Tahoma, Verdana, Segoe, sans-serif;line-height:1.2;padding-top:15px;padding-right:10px;padding-bottom:15px;padding-left:10px;"><div style="line-height: 1.2; font-size: 12px; color: #ffffff; font-family: Roboto, Tahoma, Verdana, Segoe, sans-serif; mso-line-height-alt: 14px;"><p style="font-size: 14px; line-height: 1.2; word-break: break-word; mso-line-height-alt: 17px; margin: 0;">Dapatkan akses tak terbatas baca komik di Story Tales dengan berlangganan paket Story Tales Premium. Jika ada pertanyaan seputar Story Tales, Hubungi kami kapan saja:</p></div></div><div align="center" class="button-container" style="padding-top:10px;padding-right:10px;padding-bottom:10px;padding-left:10px;"><a href="%v" style="text-decoration:none;display:inline-block;color:#575fce;background-color:#ffffff;border-radius:4px;-webkit-border-radius:4px;-moz-border-radius:4px;width:auto; width:auto;;border-top:1px solid #ffffff;border-right:1px solid #ffffff;border-bottom:1px solid #ffffff;border-left:1px solid #ffffff;padding-top:5px;padding-bottom:5px;font-family:Roboto, Tahoma, Verdana, Segoe, sans-serif;text-align:center;mso-border-alt:none;word-break:keep-all;"><span style="padding-left:20px;padding-right:20px;font-size:16px;display:inline-block;"><span style="font-size: 16px; line-height: 2; word-break: break-word; mso-line-height-alt: 32px;">%v</span></span></a></div></div></div></div></div></div></div><div style="background-color:transparent;"><div class="block-grid" style="min-width: 320px; max-width: 700px; overflow-wrap: break-word; word-wrap: break-word; word-break: break-word; Margin: 0 auto; background-color: #575fce;"><div style="border-collapse: collapse;display: table;width: 100%%;background-color:#575fce;"><div class="col num12" style="min-width: 320px; max-width: 700px; display: table-cell; vertical-align: top; width: 700px;"><div class="col_cont" style="width:100%% !important;"><div style="border-top:0px solid transparent; border-left:0px solid transparent; border-bottom:0px solid transparent; border-right:0px solid transparent; padding-top:5px; padding-bottom:5px; padding-right: 0px; padding-left: 0px;"><div></div></div></div></div></div></div></div></td></tr></tbody></table></body></html>`, t.title,
		t.logo, t.foreword, t.content,
		t.message, t.footer,
		t.logoI, t.mailLink, t.mailLabel)
}

func (t *template) Compose() string {
	t.setAsset()
	return fmt.Sprintf(`<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional //EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd"><html xmlns="http://www.w3.org/1999/xhtml" xmlns:o="urn:schemas-microsoft-com:office:office" xmlns:v="urn:schemas-microsoft-com:vml"><head><meta content="text/html; charset=utf-8" http-equiv="Content-Type"/><meta content="width=device-width" name="viewport"/><meta content="IE=edge" http-equiv="X-UA-Compatible"/><title>%v</title><style type="text/css">body{margin: 0;padding: 0;}table,td,tr{vertical-align: top;border-collapse: collapse;}*{line-height: inherit;}a[x-apple-data-detectors=true]{color: inherit !important;text-decoration: none !important;}</style><style id="media-query" type="text/css">@media (max-width: 720px){.block-grid,.col{min-width: 320px !important;max-width: 100%% !important;display: block !important;}.block-grid{width: 100%% !important;}.col{width: 100%% !important;}.col_cont{margin: 0 auto;}img.fullwidth,img.fullwidthOnMobile{max-width: 100%% !important;}.no-stack .col{min-width: 0 !important;display: table-cell !important;}.no-stack.two-up .col{width: 50%% !important;}.no-stack .col.num2{width: 16.6%% !important;}.no-stack .col.num3{width: 25%% !important;}.no-stack .col.num4{width: 33%% !important;}.no-stack .col.num5{width: 41.6%% !important;}.no-stack .col.num6{width: 50%% !important;}.no-stack .col.num7{width: 58.3%% !important;}.no-stack .col.num8{width: 66.6%% !important;}.no-stack .col.num9{width: 75%% !important;}.no-stack .col.num10{width: 83.3%% !important;}.video-block{max-width: none !important;}.mobile_hide{min-height: 0px;max-height: 0px;max-width: 0px;display: none;overflow: hidden;font-size: 0px;}.desktop_hide{display: block !important;max-height: none !important;}}</style></head><body class="clean-body" style="margin: 0; padding: 0; -webkit-text-size-adjust: 100%%; background-color: #FFFFFF;"><table bgcolor="#FFFFFF" cellpadding="0" cellspacing="0" class="nl-container" role="presentation" style="table-layout: fixed; vertical-align: top; min-width: 320px; border-spacing: 0; border-collapse: collapse; mso-table-lspace: 0pt; mso-table-rspace: 0pt; background-color: #FFFFFF; width: 100%%;" valign="top" width="100%%"><tbody><tr style="vertical-align: top;" valign="top"><td style="word-break: break-word; vertical-align: top;" valign="top"><div style="background-color:transparent;"><div class="block-grid mixed-two-up" style="min-width: 320px; max-width: 700px; overflow-wrap: break-word; word-wrap: break-word; word-break: break-word; Margin: 0 auto; background-color: transparent;"><div style="border-collapse: collapse;display: table;width: 100%%;background-color:transparent;"><div class="col num3" style="display: table-cell; vertical-align: top; max-width: 320px; min-width: 174px; width: 175px;"><div class="col_cont" style="width:100%% !important;"><div style="border-top:0px solid transparent; border-left:0px solid transparent; border-bottom:0px solid transparent; border-right:0px solid transparent; padding-top:5px; padding-bottom:5px; padding-right: 0px; padding-left: 0px;"><div align="left" class="img-container left autowidth" style="padding-right: 0px;padding-left: 10px;"><div style="font-size:1px;line-height:5px"></div><img alt="Digisoul Logo" border="0" class="left autowidth" src="%v" style="text-decoration: none; -ms-interpolation-mode: bicubic; height: auto; border: 0; width: 100%%; max-width: 117px; display: block;" title="Alternate text" width="117"/><div style="font-size:1px;line-height:5px"></div></div></div></div></div><div class="col num9" style="display: table-cell; vertical-align: top; max-width: 320px; min-width: 522px; width: 525px;"><div class="col_cont" style="width:100%% !important;"><div style="border-top:0px solid transparent; border-left:0px solid transparent; border-bottom:0px solid transparent; border-right:0px solid transparent; padding-top:5px; padding-bottom:5px; padding-right: 0px; padding-left: 0px;"><div></div></div></div></div></div></div></div><div style="background-color:transparent;"><div class="block-grid" style="min-width: 320px; max-width: 700px; overflow-wrap: break-word; word-wrap: break-word; word-break: break-word; Margin: 0 auto; background-color: transparent;"><div style="border-collapse: collapse;display: table;width: 100%%;background-color:transparent;"><div class="col num12" style="min-width: 320px; max-width: 700px; display: table-cell; vertical-align: top; width: 700px;"><div class="col_cont" style="width:100%% !important;"><div style="border-top:0px solid transparent; border-left:0px solid transparent; border-bottom:0px solid transparent; border-right:0px solid transparent; padding-top:5px; padding-bottom:5px; padding-right: 0px; padding-left: 0px;"><div style="color:#555555;font-family:Roboto, Tahoma, Verdana, Segoe, sans-serif;line-height:1.2;padding-top:15px;padding-right:10px;padding-bottom:15px;padding-left:10px;"><div style="line-height: 1.2; font-size: 12px; color: #555555; font-family: Roboto, Tahoma, Verdana, Segoe, sans-serif; mso-line-height-alt: 14px;"><p style="font-size: 14px; line-height: 1.2; word-break: break-word; mso-line-height-alt: 17px; margin: 0;"><strong>%v</strong></p></div></div></div></div></div></div></div></div><div style="background-color:transparent;"><div class="block-grid" style="min-width: 320px; max-width: 700px; overflow-wrap: break-word; word-wrap: break-word; word-break: break-word; Margin: 0 auto; background-color: transparent;"><div style="border-collapse: collapse;display: table;width: 100%%;background-color:transparent;"><div class="col num12" style="min-width: 320px; max-width: 700px; display: table-cell; vertical-align: top; width: 700px;"><div class="col_cont" style="width:100%% !important;"><div style="border-top:0px solid transparent; border-left:0px solid transparent; border-bottom:0px solid transparent; border-right:0px solid transparent; padding-top:5px; padding-bottom:5px; padding-right: 0px; padding-left: 0px;"><div style="color:#555555;font-family:Roboto, Tahoma, Verdana, Segoe, sans-serif;line-height:1.2;padding-top:10px;padding-right:10px;padding-bottom:10px;padding-left:10px;"><div style="line-height: 1.2; font-size: 12px; color: #555555; font-family: Roboto, Tahoma, Verdana, Segoe, sans-serif; mso-line-height-alt: 14px;"><p style="font-size: 14px; line-height: 1.2; word-break: break-word; mso-line-height-alt: 17px; margin: 0;">%v</p></div></div></div></div></div></div></div></div><div style="background-color:transparent;"><div class="block-grid" style="min-width: 320px; max-width: 700px; overflow-wrap: break-word; word-wrap: break-word; word-break: break-word; Margin: 0 auto; background-color: transparent;"><div style="border-collapse: collapse;display: table;width: 100%%;background-color:transparent;"><div class="col num12" style="min-width: 320px; max-width: 700px; display: table-cell; vertical-align: top; width: 700px;"><div class="col_cont" style="width:100%% !important;"><div style="border-top:0px solid transparent; border-left:0px solid transparent; border-bottom:0px solid transparent; border-right:0px solid transparent; padding-top:5px; padding-bottom:5px; padding-right: 0px; padding-left: 0px;"><div align="center" class="button-container" style="padding-top:10px;padding-right:10px;padding-bottom:10px;padding-left:10px;"><a href="%v" style="-webkit-text-size-adjust: none; text-decoration: none; display: inline-block; color: #ffffff; background-color: #575fce; border-radius: 4px; -webkit-border-radius: 4px; -moz-border-radius: 4px; width: auto; width: auto; border-top: 1px solid #575fce; border-right: 1px solid #575fce; border-bottom: 1px solid #575fce; border-left: 1px solid #575fce; padding-top: 5px; padding-bottom: 5px; font-family: Roboto, Tahoma, Verdana, Segoe, sans-serif; text-align: center; mso-border-alt: none; word-break: keep-all;" target="_blank"><span style="padding-left:20px;padding-right:20px;font-size:16px;display:inline-block;"><span style="font-size: 16px; line-height: 2; word-break: break-word; mso-line-height-alt: 32px;">%v</span></span></a></div></div></div></div></div></div></div><div style="background-color:transparent;"><div class="block-grid" style="min-width: 320px; max-width: 700px; overflow-wrap: break-word; word-wrap: break-word; word-break: break-word; Margin: 0 auto; background-color: transparent;"><div style="border-collapse: collapse;display: table;width: 100%%;background-color:transparent;"><div class="col num12" style="min-width: 320px; max-width: 700px; display: table-cell; vertical-align: top; width: 700px;"><div class="col_cont" style="width:100%% !important;"><div style="border-top:0px solid transparent; border-left:0px solid transparent; border-bottom:0px solid transparent; border-right:0px solid transparent; padding-top:5px; padding-bottom:5px; padding-right: 0px; padding-left: 0px;"><div style="color:#555555;font-family:Roboto, Tahoma, Verdana, Segoe, sans-serif;line-height:1.2;padding-top:10px;padding-right:10px;padding-bottom:15px;padding-left:10px;"><div style="line-height: 1.2; font-size: 12px; color: #555555; font-family: Roboto, Tahoma, Verdana, Segoe, sans-serif; mso-line-height-alt: 14px;"><p style="font-size: 14px; line-height: 1.2; word-break: break-word; mso-line-height-alt: 17px; margin: 0;">%v</p></div></div></div></div></div></div></div></div><div style="background-color:transparent;"><div class="block-grid mixed-two-up" style="min-width: 320px; max-width: 700px; overflow-wrap: break-word; word-wrap: break-word; word-break: break-word; Margin: 0 auto; background-color: transparent;"><div style="border-collapse: collapse;display: table;width: 100%%;background-color:transparent;"><div class="col num4" style="display: table-cell; vertical-align: top; max-width: 320px; min-width: 232px; width: 233px;"><div class="col_cont" style="width:100%% !important;"><div style="border-top:0px solid transparent; border-left:0px solid transparent; border-bottom:0px solid transparent; border-right:0px solid transparent; padding-top:5px; padding-bottom:5px; padding-right: 0px; padding-left: 0px;"><div style="color:#555555;font-family:Roboto, Tahoma, Verdana, Segoe, sans-serif;line-height:1.2;padding-top:10px;padding-right:10px;padding-bottom:25px;padding-left:10px;"><div style="line-height: 1.2; font-size: 12px; color: #555555; font-family: Roboto, Tahoma, Verdana, Segoe, sans-serif; mso-line-height-alt: 14px;"><p style="font-size: 14px; line-height: 1.2; word-break: break-word; mso-line-height-alt: 17px; margin: 0;">Salam,</p></div></div><div style="color:#555555;font-family:Roboto, Tahoma, Verdana, Segoe, sans-serif;line-height:1.2;padding-top:30px;padding-right:10px;padding-bottom:15px;padding-left:10px;"><div style="line-height: 1.2; font-size: 12px; color: #555555; font-family: Roboto, Tahoma, Verdana, Segoe, sans-serif; mso-line-height-alt: 14px;"><p style="line-height: 1.2; word-break: break-word; font-size: 14px; mso-line-height-alt: 17px; margin: 0;"><span style="font-size: 14px;">Tim StoryTale</span></p></div></div></div></div></div><div class="col num8" style="display: table-cell; vertical-align: top; max-width: 320px; min-width: 464px; width: 466px;"><div class="col_cont" style="width:100%% !important;"><div style="border-top:0px solid transparent; border-left:0px solid transparent; border-bottom:0px solid transparent; border-right:0px solid transparent; padding-top:5px; padding-bottom:5px; padding-right: 0px; padding-left: 0px;"><div></div></div></div></div></div></div></div><div style="background-color:transparent;"><div class="block-grid mixed-two-up" style="min-width: 320px; max-width: 700px; overflow-wrap: break-word; word-wrap: break-word; word-break: break-word; Margin: 0 auto; background-color: #575fce;"><div style="border-collapse: collapse;display: table;width: 100%%;background-color:#575fce;"><div class="col num4" style="display: table-cell; vertical-align: top; max-width: 320px; min-width: 232px; width: 233px;"><div class="col_cont" style="width:100%% !important;"><div style="border-top:0px solid transparent; border-left:0px solid transparent; border-bottom:0px solid transparent; border-right:0px solid transparent; padding-top:5px; padding-bottom:5px; padding-right: 0px; padding-left: 0px;"><div align="left" class="img-container left autowidth" style="padding-right: 0px;padding-left: 15px;"><div style="font-size:1px;line-height:5px"></div><img alt="Alternate text" border="0" class="left autowidth" src="%v" style="text-decoration: none; -ms-interpolation-mode: bicubic; height: auto; border: 0; width: 100%%; max-width: 116px; display: block;" title="Alternate text" width="116"/></div></div></div></div><div class="col num8" style="display: table-cell; vertical-align: top; max-width: 320px; min-width: 464px; width: 466px;"><div class="col_cont" style="width:100%% !important;"><div style="border-top:0px solid transparent; border-left:0px solid transparent; border-bottom:0px solid transparent; border-right:0px solid transparent; padding-top:5px; padding-bottom:5px; padding-right: 0px; padding-left: 0px;"><div></div></div></div></div></div></div></div><div style="background-color:transparent;"><div class="block-grid" style="min-width: 320px; max-width: 700px; overflow-wrap: break-word; word-wrap: break-word; word-break: break-word; Margin: 0 auto; background-color: #575fce;"><div style="border-collapse: collapse;display: table;width: 100%%;background-color:#575fce;"><div class="col num12" style="min-width: 320px; max-width: 700px; display: table-cell; vertical-align: top; width: 700px;"><div class="col_cont" style="width:100%% !important;"><div style="border-top:0px solid transparent; border-left:0px solid transparent; border-bottom:0px solid transparent; border-right:0px solid transparent; padding-top:5px; padding-bottom:5px; padding-right: 0px; padding-left: 0px;"><div style="color:#ffffff;font-family:Roboto, Tahoma, Verdana, Segoe, sans-serif;line-height:1.2;padding-top:15px;padding-right:10px;padding-bottom:15px;padding-left:10px;"><div style="line-height: 1.2; font-size: 12px; color: #ffffff; font-family: Roboto, Tahoma, Verdana, Segoe, sans-serif; mso-line-height-alt: 14px;"><p style="font-size: 14px; line-height: 1.2; word-break: break-word; mso-line-height-alt: 17px; margin: 0;">Dapatkan akses tak terbatas baca komik di Story Tales dengan berlangganan paket Story Tales Premium. Jika ada pertanyaan seputar Story Tales, Hubungi kami kapan saja:</p></div></div><div align="center" class="button-container" style="padding-top:10px;padding-right:10px;padding-bottom:10px;padding-left:10px;"><a href="%v" style="text-decoration:none;display:inline-block;color:#575fce;background-color:#ffffff;border-radius:4px;-webkit-border-radius:4px;-moz-border-radius:4px;width:auto; width:auto;;border-top:1px solid #ffffff;border-right:1px solid #ffffff;border-bottom:1px solid #ffffff;border-left:1px solid #ffffff;padding-top:5px;padding-bottom:5px;font-family:Roboto, Tahoma, Verdana, Segoe, sans-serif;text-align:center;mso-border-alt:none;word-break:keep-all;"><span style="padding-left:20px;padding-right:20px;font-size:16px;display:inline-block;"><span style="font-size: 16px; line-height: 2; word-break: break-word; mso-line-height-alt: 32px;">%v</span></span></a></div></div></div></div></div></div></div><div style="background-color:transparent;"><div class="block-grid" style="min-width: 320px; max-width: 700px; overflow-wrap: break-word; word-wrap: break-word; word-break: break-word; Margin: 0 auto; background-color: #575fce;"><div style="border-collapse: collapse;display: table;width: 100%%;background-color:#575fce;"><div class="col num12" style="min-width: 320px; max-width: 700px; display: table-cell; vertical-align: top; width: 700px;"><div class="col_cont" style="width:100%% !important;"><div style="border-top:0px solid transparent; border-left:0px solid transparent; border-bottom:0px solid transparent; border-right:0px solid transparent; padding-top:5px; padding-bottom:5px; padding-right: 0px; padding-left: 0px;"><div></div></div></div></div></div></div></div></td></tr></tbody></table></body></html>`,
		t.title,
		t.logo, t.foreword, t.content,
		t.buttonLink, t.buttonLabel, t.footer,
		t.logoI, t.mailLink, t.mailLabel)
}
