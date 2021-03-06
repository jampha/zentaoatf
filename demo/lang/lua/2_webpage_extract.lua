#!/usr/bin/env lua
--[[
[case]
title=extract content from webpage
cid=0
pid=0

[group]
  1. Load web page from url http://xxx 
  2. Retrieve img element zt-logo.png in html 
  3. Check img exist >> .*zt-logo.png

[esac]
]]

local http = require("socket.http") -- need luasocket library (luarocks install luasocket)
local ltn12 = require("ltn12")

function http.get(u)
   local t = {}
   local r, c, h = http.request{
      url = u,
      sink = ltn12.sink.table(t)}
   return r, c, h, table.concat(t)
end

r,c,h,body = http.get("http://pms.zentao.net/user-login.html")
if c~= 200 then
    print("ERR: " .. c)
else
    _, _, src = string.find(body, "<img%ssrc='(.-)' .*>")
    print(">>" .. src)
end