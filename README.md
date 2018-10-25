# Handbook Generator

**This project is no longer maintained by Turbine Labs, which has
[shut down](https://blog.turbinelabs.io/turbine-labs-is-shutting-down-and-our-team-is-joining-slack-2ad41554920c).**

This tool is designed to templatize the Clef employee handbook, allowing other companies to use it as a reasonable starting point without excessive amounts of s/Clef/company name/. The goals in development were to

* enable a smooth migration of the template system back into the Clef handbook github repo
* allow other companies to easily adopt the current Clef handbook "as is"
* allow other companies to easily merge changes as they are made to the Clef handbook
* allow other companies to easily replace sections of the handbook as they desire

# Basic usage

This assumes you have the go compiler and tools installed. See https://golang.org/doc/install#install for instructions. From this directory, you should be able to run `go run generator.go`. This should effectively re-create the Clef handbook in the a directory named "out". Running `go run generate.go --help` will provide a description of all command line arguments available.

# Customizing the handbook

Customization is done via two mechanisms - overrides and vars. Vars are held (by default) in the vars.json file. Vars are interpolated into the handbook templates, held in the root of this git repository. Editing the vars file and re-running the generator will create a modified handobok in the out directory. Modifying vars is useful for changes like company name, founder emails, etc.

Overrides allow forks to replace sections of the handbook entirely while still easily tracking changes from an upstream. When selecting a template, the generator walks the directory specified by the `-template-dir` flag (by default templates). When it encounters a file, it looks for a matching file in overrides. If such a file exists, it uses the override as the template. For instance when the generator encounters `templates/Employment Policies/Salary and Equity Compensation.md`, it will look for `overrides/Employment Policies/Salary and Equity Compensation.md`. If the override file is found it will be used as the template for that page in the generated handbook.

# Suggested Usage

## Managing Updates

The Clef handbook is a living document, and changes have been made fairly rapidly. We don't want to stifle that rapid evolution, but we would like to make it possible for more people to share in the benefits. Our suggested workflow is to have two repositories - a handbook-template and an actual handbook. The handbook-template repository wouldbe shared between multiple companies, and treated much more like a standard open source project than the current purely textual Clef handbook. Multiple companies could collaborate here, share changes, and manage forks to adapt the handbook to their specific needs. When the handbook-template is ready for publishing, the generator is run and output is written to the actual handbook repository. As an example

```
*company forks clef/handbook-template*
*company creates <company>/handbook repo*
git clone git@github.com:<company>/handbook-template.git
git clone git@github.com:<company>/handbook.git
cd handbook-template
go run generator.go -out ../handbook
cd ../handbook
git commit -am "initial commit" && git push origin master
*make edits to handbook-template*
*pull changes from clef/handbook-template*
cd handbook-template
go run generator.go -out ../handbook
cd ../handbook
git commit -am "updates" && git push origin master
```

## Probable Overrides

The following areas will probably see fairly wide divergence across companies. This list may seem daunting, but the override process is straightforward. Copy templates you want to change to the overrides file, make your changes, and regenerate your handbook.

* Policy Changes.md - a very detailed list of change policies, but once again Clef specific. A good starting point, but you'll want to override
* Values.md - these are clef values. You should have your own!
* Benefits and Perks/Healthcare and Disability Insurance.md - companies are almost certain to have different providers and levels of healthcare. You'll probabyl want to override this file
* Employment Policies/Working Remotely.md - a hot topic, if you support remote work you'll want to override this file as well.
* Hiring Documents/Guide to Your Equity.md - this contains some fairly Clef-specific language, and will likely be an override a well.
* Onboarding Documents/Product Manifesto.md - this is entirely Clef-specific. Make your own manifesto!
* Onboarding Documents/Welcome.md - also contains some fairly Clef-specific language.
* Operations Documents/Onboarding.md - Slack and Trello are common but probably not _standard_. You may want to override here as well.
* Operations Documents/Sharing Files.md - This provides a fairly straightforward way to use Google Drive, but your mileage may vary here
* Operations Documents/Sourcing Candidates.md - This contains a set of specific organizationts Clef works with to source candidates. If this is your list too, great! If not, you'll want to override.

