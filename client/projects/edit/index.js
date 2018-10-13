import Overall from './Overall'
import Story from './Story'
import Promises from './Promises'
import Finances from './Finances'

export default [{
	path: '',
	name: 'projectEditOverall',
	component: Overall,
	componentTitle: "Get started",
	componentDescription: "Make the overall decisions",
	pageName: "Overall",
}, {
	path: 'story',
	name: 'projectEditStory',
	component: Story,
	componentTitle: "Tell the story",
	componentDescription: "Make the overall decisions",
	pageName: "Story",
}, {
	path: 'promises',
	name: 'projectEditPromises',
	component: Promises,
	componentTitle: "Get started",
	componentDescription: "Make the overall decisions",
	pageName: "Promises",
}, {
	path: 'finances',
	name: 'projectEditFinances',
	component: Finances,
	componentTitle: "Get started",
	componentDescription: "Make the overall decisions",
	pageName: "Finances",
}]
