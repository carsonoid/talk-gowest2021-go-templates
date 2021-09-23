Go ships with with the text/template and html/template packages built-in! Both of these work exceptionally well. In fact the text/template system is powerful enough to be the core component for the Kubernetes package manager Helm.

So why do people often dismiss Go templates? There are extensive GoDocs for the text/template package but some very key concepts often go overlooked. If you have ever been confused by automatic scoping or errors about an incorrect number of args then you have probably fallen victim to some very early stumbling blocks.

However, anyone that can take a moment and master some core concepts that come with the template language will find that it is actually very extensible and robust. This talk is here to help you get past the common early blocks and take your template usage to the next level.

Come watch this talk to see how you can harness these packages and take your templates from this:

```
Go templates are {{ index .Emojis "fire" }}! {{ if not (eq .ScaryThing "") }} Just don't get scared by the "{{ .ScaryThing }}" {{ end }}
```

To this:

```
Go templates are {{ emoji "fire" }}! {{- with .ScaryThing }} Just don't get scared by the {{ . | quote }} {{- end }}
```

---

I do a lot of work with go templates nearly every day in Go. I have also done a lot of deep dives on other template languages like Jinja as part of my work with SaltStack.

Given that experience, I feel that Go templates are often too easily dismissed by those that have not taken the time to truly understand them. There are some early concepts that often give people a really bad first impression. For example: automatic scoping tends to really confuse and frustrate users early on.

I also feel like very few users take full advantage of the extensibility options included with templates. Things like using FuncMaps + Closures to create powerful and easy to use additions to templates. Or users not understanding how to tuple/index multiple pieces of data into and out of sub-templates. Not to mention that with the recent addition of go:embed we can now start keeping our templates in dedicated files instead of big raw string constants!

Templating in a compiled and strictly-typed language is never going to be quite as robust as with a dynamically-typed scripting language but I feel that we often don’t respect just how good the built-in go template system is when used to it’s full extent.

