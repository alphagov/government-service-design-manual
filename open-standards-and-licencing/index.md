---
layout: default
title: Open standards and licensing
subtitle: Building services using common standards
section: guidance
type: guide
audience:
  primary: tech-arch, developer
  secondary: service-manager
status: draft
theme: design-and-build
---
    
#Open standards and licensing
Herein, best practice and considerations for using open standards and using, publishing and contributing to free and open source software ([FOSS](http://en.wikipedia.org/wiki/Free_and_open-source_software)).
Please note, this is for guidance purposes only, and should not be taken as [legal advice](http://en.wikipedia.org/wiki/IANACL).

##Guidance
Preference should be given to using open standards with the broadest remit over local standards
as prescribed in the [consultation on open standards](http://consultation.cabinetoffice.gov.uk/openstandards/)

As with any software, using open source comes with an obligation to follow the terms of the software licence.
FOSS differs from [proprietary software](http://en.wikipedia.org/wiki/Proprietary_software) in that the source code may be modified,
in which case the licence may demand any derived code be republished with [attribution](http://en.wikipedia.org/wiki/Attribution_(copyright),
and under the same terms of the original licence, sometimes called [copyleft](http://en.wikipedia.org/wiki/Copyleft).
These obligations may introduce difficulties when mixing open source software with closed source software, particularly
when combining [GPL](http://en.wikipedia.org/wiki/GNU_General_Public_License) software with software developed behind closed doors
or using [GNU AFFERO](http://www.gnu.org/licenses/agpl-3.0.html) licensed software in proprietary services.
These risks may be mitigated by [developing software in the open](http://digital.cabinetoffice.gov.uk/2012/10/12/coding-in-the-open/).

Software should be published as open source, under a [permissive free software licence](http://en.wikipedia.org/wiki/Permissive_free_software_licence) such as [BSD](http://opensource.org/licenses/BSD-2-Clause) or [MIT](http://opensource.org/licenses/MIT).

##Templates
The following text is the [MIT licence](http://opensource.org/licenses/MIT) used by GDS when publishing software in the open:
<pre>
Copyright (C) 2012 HM Government (Government Digital Service)

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
of the Software, and to permit persons to whom the Software is furnished to do
so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
</pre>
It is sufficient to include this as a single **LICENSE** [sic] file, though you may elect to include the licence in 
individual source code files, particularly those which may be taken in isolation, such as example code.

##Why we do this
There are many other advantages to [developing software in the open](http://digital.cabinetoffice.gov.uk/2012/10/12/coding-in-the-open/)
described in more detail in the [Technical Architecture/Open source considerations]() section,
and prescribed by the tenth GDS design principle: [Make things open: it makes things better](https://www.gov.uk/designprinciples#tenth).

##Further reading
- [fossbazaar.org](https://fossbazaar.org/) is an interest group for the governance of FOSS in organizations.
- the [International Free and Open Source Software Law Review](http://www.ifosslr.org/) is a collaborative legal publication aiming to increase knowledge and understanding among lawyers about Free and Open Source Software issues. Topics covered include copyright, licence implementation, licence interpretation, software patents, open standards, case law and statutory changes.
- [fossology](http://www.fossology.org/) is one of a number of open source tools useful when assessing open source used by a project.
