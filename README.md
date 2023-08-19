# 青年大学习skip

[![Netlify Status](https://api.netlify.com/api/v1/badges/93a7e121-af8d-46d0-9019-32b69ba447d6/deploy-status)](https://app.netlify.com/sites/qndxx/deploys)
[![Vercel Status](https://therealsujitk-vercel-badge.vercel.app/?app=vercel.com/hx-w/qndxx)](https://vercel.com/hx-w/qndxx)

**该项目旨在应付每周学生的青年大学习检查，以便节省你一周大约10分钟的时间和1%的手机电量。**

**此方法并不会在青年大学习后台留下观看记录，谨慎使用。**

> 可靠的做法是点开大学习后立马关闭，在后台就已经留有观看记录，如果需要的话再通过该库获得截图。

## 快速开始 | Quick Start

1. 打开手机微信
2. 在微信中点开指定链接
3. 截图

> 站点部署在netlify平台上，第一次访问速度会较慢请耐心等待。

## 站点 | Website

目前暂时需要手动在url参数中输入最新一期的标题，相关进度跟进见 [Issue#13](https://github.com/hx-w/qndxx/issues/13)

- [x] [**qndxx.scubot.com?id=2023年第18期**](https://qndxx.scubot.com?id=2023年第18期)
- [ ] [**qndxx.netlify.app**](https://qndxx.netlify.app)
- [ ] [**qndxx-skip.vercel.app**](https://qndxx-skip.vercel.app)
- [ ] [**dxx.haha44444.top**](https://dxx.haha44444.top)
- [ ] [**dxx.txt1482.ltd**](https://dxx.txt1482.ltd)

> 微信可能封禁部分域名，不定期更新有效站点。


## 部署 | Deploy

如果你想自己部署一个该项目（~虽然没啥必要~），那么我建议从以下几个方式选择。

**Vercel**

[![Deploy with Vercel](https://vercel.com/button)](https://vercel.com/new/clone?repository-url=https%3A%2F%2Fgithub.com%2Fhx-w%2Fqndxx&project-name=qndxx&repository-name=qndxx)

1. Vercel的管理站点与部署后的站点在国内都不能直连，确定继续的话就点击上面的按钮
2. 将 `Settings >> Git >> Production Branch` 设为 `prod-vercel`
3. 将 `Settings >> General >> Node.js Version` 设为 `16.x`
4. 在 `Deployments` 下选取最近的部署记录，进行Redeploy

**Netlify**

[![Deploy to Netlify](https://www.netlify.com/img/deploy/button.svg)](https://app.netlify.com/start/deploy?repository=https://github.com/hx-w/qndxx)
