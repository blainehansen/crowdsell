<template lang="pug">

#how-it-works.container
	//- future ideas:
	//- crowdfund investment for ideas that aren't up
	//- torrent hosting

	h1 How Does Crowdsell Work Exactly?

	p Crowdsell is all about making crowdfunding rigorous and dependable.

	p Here are the exact steps that each project goes through, in exact detail.

	ul
		li
			| The creator puts together the project on the site. They write the title and the description, choose a funding amount, and give their banking information where they will receive the raise money.
		li
			| Then they upload the very important
			i preview material
			| to showcase to the whole world what exactly their project is, why it's worth contributing to, and to give a taste of the quality of the work. They also make a list of promises about the work, which are expected to be objective true or false statements about the work being offered. After the project closes, the work will be evaluated against these promises. These are just meant to act as a way to assure the community of what they're being offered.
			p It's important to note that at this point, the work is still basically a secret, and only the creator and their team has access to it.

	quote
		p
			b Example: An App
		p The preview material might be screencasts of the app being used, screenshots of every page, a "lite" version with some important features removed, and ideally the full suite of software unit tests.
		p The promises could be a list of every feature the app has, a guarantee that all the provided unit tests pass, and perhaps even a description of the full file tree of the app's source code.

	quote
		p
			b Example: A New Video Game
		p The preview material for a video game would likely be any trailers and screenshots, and probably a short demo or the first few minutes of the game.
		p The promises for the quality of a work of art will necessarily be a lot less comprehensive than something technical. They might be the system and specs the game requires, an average length based on playtesting, a list of features, and perhaps a list of the levels or areas.

	p
		| Visit
		router-link(to="/who-is-crowdsell-for") the page about who and what Crowdsell is for
		| to see more examples.

	ul
		li With all of the project information ready to be shown to the world, the creator submits the project to be reviewed by Crowdsell (to make sure it isn't doing any illegal or horrifying).
		li Once it's approved, it's published on the site, ready for contributions.

		li
			| On Crowdsell, when someone pledges to a project, they are charged immediately and the money goes into escrow. To read more about this,
			router-link(to="/why-is-crowdsell-different") visit the page about Crowdsell's differences from other crowdfunding systems.

		li
			| For a short period at the beginning of the campaign (called the "blind period"), no one, not even the project creators, will know exactly how much the project has raised. To understand why we made this decision,
			router-link(to="/why-is-crowdsell-different") visit the page about Crowdsell's differences from other crowdfunding systems.

		li After the blind period, the project will stay open until one these things happens:
			ul
				li The project raises its full amount. If this happens, the project will close and the approval and creator payment will begin.
				li The creator decides that they are fine with the amount already raised, and want to just accept it and close the project.

	quote
		p
			b Example: a project that surpasses its goal during the blind period.

		p While the blind period is going, the community can go above the project goal without the project closing. So for a project that does surpass its goal, once the blind period ends and the raised amount is made public, the project will successfully close immediately.

	quote
		p
			b Example: a project that doesn't reach its goal during the blind period.

		p This will probably be the most common case. We encourage creators to set goals that properly compensate them for their hard work, and that can fund future work on the project or other exciting projects, so it might take a while to reach the goal. The second it does reach the goal though, the project successfully closes.

	ul
		li After the project closes, the creator releases the work, and it's now the backers' job to evaluate the project and make sure it lives up to the promises the creator made.

		li Every backer will be sent an email asking them to vote on if the project lived up to its promises. They can give three basic answers:
			ul
				li "Yes, everything looks great!" This is a vote to move forward and give the creator the money.
				li "No, some promises are broken." This is a vote to not move forward and not give the creator the money. This is a serious vote, we'll talk in a minute about how it will be handled.
				li
					| "Yes, but I have some criticisms." This is a vote to move forward, but the backer wants to make public some problem they have. This vote is for situations where the project
					i technically
					| lives up to its promises, but the backer feels the work still doesn't quite live up. This will probably mostly happen for artistic projects where the backer feels that the creator didn't lie, but also didn't try their best. These criticisms will be published to the project page, along with what percentage the person contributed to the project goal. They can choose to remain anonymous.

		li The backers have one week (exactly 168 hours from the time the project closed) to make their votes. Any person who hasn't voted by that time will simply be considered a "Yes".

		li Each backer is given an amount of votes according to how much they pledged. This is just to make sure that some person or group doesn't try to cheat the system by making many small pledges so they can "stuff the ballot box".

		li After the voting period closes, the entire vote catalog will be anonymized and published so that each person who contributed can check that their vote was counted correctly. We'll talk more in another article about this cryptographic process.

		li If five percent or more of the projects votes come in as "No", the project will be reviewed by someone at Crowdsell. That person will work with the community to see if they can either resolve the situation (by allowing the creator to fix the problems and let the community approve again), or to simply condemn the project and refund all the backers. This latter situation is a sad one, since the creator who lied will be permanently banned from the site, and all the backers will get their money back at a roughly ten percent loss because of all the charge fees. But defrauding people is serious, and Crowdsell is all about making crowdfunding rigorous and dependable.

		li But, if the votes come back positive, then the project funds will be paid to the creator and everyone can celebrate a success! The commons is richer now, and the whole world has a free and open new work to enjoy.


</template>
